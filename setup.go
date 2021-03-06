/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Community License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Community-1.0.0.md

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by Kubeform. DO NOT EDIT.

package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync"
	"time"

	ec "github.com/elastic/terraform-provider-ec/ec"
	"github.com/gobuffalo/flect"
	auditlib "go.bytebuilders.dev/audit/lib"
	arv1 "k8s.io/api/admissionregistration/v1"
	"k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	informers "k8s.io/apiextensions-apiserver/pkg/client/informers/externalversions"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	admissionregistrationv1 "k8s.io/client-go/kubernetes/typed/admissionregistration/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/klog/v2"
	deploymentv1alpha1 "kubeform.dev/provider-ec-api/apis/deployment/v1alpha1"
	controllersdeployment "kubeform.dev/provider-ec-controller/controllers/deployment"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

var _provider = ec.Provider()

var runningControllers = struct {
	sync.RWMutex
	mp map[schema.GroupVersionKind]bool
}{mp: make(map[schema.GroupVersionKind]bool)}

func watchCRD(ctx context.Context, crdClient *clientset.Clientset, vwcClient *admissionregistrationv1.AdmissionregistrationV1Client, stopCh <-chan struct{}, mgr manager.Manager, auditor *auditlib.EventPublisher, restrictToNamespace string) error {
	informerFactory := informers.NewSharedInformerFactory(crdClient, time.Second*30)
	i := informerFactory.Apiextensions().V1().CustomResourceDefinitions().Informer()
	l := informerFactory.Apiextensions().V1().CustomResourceDefinitions().Lister()

	i.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			var key string
			key, err := cache.MetaNamespaceKeyFunc(obj)
			if err != nil {
				klog.Error(err)
				return
			}

			_, name, err := cache.SplitMetaNamespaceKey(key)
			if err != nil {
				klog.Error(err)
				return
			}

			crd, err := l.Get(name)
			if err != nil {
				klog.Error(err)
				return
			}
			if strings.Contains(crd.Spec.Group, "ec.kubeform.com") {
				gvk := schema.GroupVersionKind{
					Group:   crd.Spec.Group,
					Version: crd.Spec.Versions[0].Name,
					Kind:    crd.Spec.Names.Kind,
				}

				// check whether this gvk came before, if no then start the controller
				runningControllers.RLock()
				_, ok := runningControllers.mp[gvk]
				runningControllers.RUnlock()

				if !ok {
					runningControllers.Lock()
					runningControllers.mp[gvk] = true
					runningControllers.Unlock()

					if enableValidatingWebhook {
						// add dynamic ValidatingWebhookConfiguration

						// create empty VWC if the group has come for the first time
						err := createEmptyVWC(vwcClient, gvk)
						if err != nil {
							klog.Error(err)
							return
						}

						// update
						err = updateVWC(vwcClient, gvk)
						if err != nil {
							klog.Error(err)
							return
						}

						err = SetupWebhook(mgr, gvk)
						if err != nil {
							setupLog.Error(err, "unable to enable webhook")
							os.Exit(1)
						}
					}

					err = SetupManager(ctx, mgr, gvk, auditor, restrictToNamespace)
					if err != nil {
						setupLog.Error(err, "unable to start manager")
						os.Exit(1)
					}
				}
			}
		},
	})

	informerFactory.Start(stopCh)

	return nil
}

func createEmptyVWC(vwcClient *admissionregistrationv1.AdmissionregistrationV1Client, gvk schema.GroupVersionKind) error {
	vwcName := strings.ReplaceAll(strings.ToLower(gvk.Group), ".", "-")
	_, err := vwcClient.ValidatingWebhookConfigurations().Get(context.TODO(), vwcName, metav1.GetOptions{})
	if err == nil || !(errors.IsNotFound(err)) {
		return err
	}

	emptyVWC := &arv1.ValidatingWebhookConfiguration{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ValidatingWebhookConfiguration",
			APIVersion: "admissionregistration.k8s.io/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: strings.ReplaceAll(strings.ToLower(gvk.Group), ".", "-"),
			Labels: map[string]string{
				"app.kubernetes.io/instance": "ec.kubeform.com",
				"app.kubernetes.io/part-of":  "kubeform.com",
			},
		},
	}
	_, err = vwcClient.ValidatingWebhookConfigurations().Create(context.TODO(), emptyVWC, metav1.CreateOptions{})
	if err != nil {
		return err
	}

	return nil
}

func updateVWC(vwcClient *admissionregistrationv1.AdmissionregistrationV1Client, gvk schema.GroupVersionKind) error {
	vwcName := strings.ReplaceAll(strings.ToLower(gvk.Group), ".", "-")
	vwc, err := vwcClient.ValidatingWebhookConfigurations().Get(context.TODO(), vwcName, metav1.GetOptions{})
	if err != nil {
		return err
	}

	path := "/validate-" + strings.ReplaceAll(strings.ToLower(gvk.Group), ".", "-") + "-v1alpha1-" + strings.ToLower(gvk.Kind)
	fail := arv1.Fail
	sideEffects := arv1.SideEffectClassNone
	admissionReviewVersions := []string{"v1beta1"}

	rules := []arv1.RuleWithOperations{
		{
			Operations: []arv1.OperationType{
				arv1.Delete,
				arv1.Update,
			},
			Rule: arv1.Rule{
				APIGroups:   []string{strings.ToLower(gvk.Group)},
				APIVersions: []string{gvk.Version},
				Resources:   []string{strings.ToLower(flect.Pluralize(gvk.Kind))},
			},
		},
	}

	data, err := ioutil.ReadFile("/tmp/k8s-webhook-server/serving-certs/ca.crt")
	if err != nil {
		return err
	}

	name := strings.ToLower(gvk.Kind) + "." + gvk.Group
	for _, webhook := range vwc.Webhooks {
		if webhook.Name == name {
			return nil
		}
	}

	newWebhook := arv1.ValidatingWebhook{
		Name: name,
		ClientConfig: arv1.WebhookClientConfig{
			Service: &arv1.ServiceReference{
				Namespace: webhookNamespace,
				Name:      webhookName,
				Path:      &path,
			},
			CABundle: data,
		},
		Rules:                   rules,
		FailurePolicy:           &fail,
		SideEffects:             &sideEffects,
		AdmissionReviewVersions: admissionReviewVersions,
	}

	vwc.Webhooks = append(vwc.Webhooks, newWebhook)

	_, err = vwcClient.ValidatingWebhookConfigurations().Update(context.TODO(), vwc, metav1.UpdateOptions{})
	if err != nil {
		return err
	}

	return nil
}

func SetupManager(ctx context.Context, mgr manager.Manager, gvk schema.GroupVersionKind, auditor *auditlib.EventPublisher, restrictToNamespace string) error {
	switch gvk {
	case schema.GroupVersionKind{
		Group:   "deployment.ec.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Deployment",
	}:
		if err := (&controllersdeployment.DeploymentReconciler{
			Client:   mgr.GetClient(),
			Log:      ctrl.Log.WithName("controllers").WithName("Deployment"),
			Scheme:   mgr.GetScheme(),
			Gvk:      gvk,
			Provider: _provider,
			Resource: _provider.ResourcesMap["ec_deployment"],
			TypeName: "ec_deployment",
		}).SetupWithManager(ctx, mgr, auditor, restrictToNamespace); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "Deployment")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "deployment.ec.kubeform.com",
		Version: "v1alpha1",
		Kind:    "ElasticsearchKeystore",
	}:
		if err := (&controllersdeployment.ElasticsearchKeystoreReconciler{
			Client:   mgr.GetClient(),
			Log:      ctrl.Log.WithName("controllers").WithName("ElasticsearchKeystore"),
			Scheme:   mgr.GetScheme(),
			Gvk:      gvk,
			Provider: _provider,
			Resource: _provider.ResourcesMap["ec_deployment_elasticsearch_keystore"],
			TypeName: "ec_deployment_elasticsearch_keystore",
		}).SetupWithManager(ctx, mgr, auditor, restrictToNamespace); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "ElasticsearchKeystore")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "deployment.ec.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Extension",
	}:
		if err := (&controllersdeployment.ExtensionReconciler{
			Client:   mgr.GetClient(),
			Log:      ctrl.Log.WithName("controllers").WithName("Extension"),
			Scheme:   mgr.GetScheme(),
			Gvk:      gvk,
			Provider: _provider,
			Resource: _provider.ResourcesMap["ec_deployment_extension"],
			TypeName: "ec_deployment_extension",
		}).SetupWithManager(ctx, mgr, auditor, restrictToNamespace); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "Extension")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "deployment.ec.kubeform.com",
		Version: "v1alpha1",
		Kind:    "TrafficFilter",
	}:
		if err := (&controllersdeployment.TrafficFilterReconciler{
			Client:   mgr.GetClient(),
			Log:      ctrl.Log.WithName("controllers").WithName("TrafficFilter"),
			Scheme:   mgr.GetScheme(),
			Gvk:      gvk,
			Provider: _provider,
			Resource: _provider.ResourcesMap["ec_deployment_traffic_filter"],
			TypeName: "ec_deployment_traffic_filter",
		}).SetupWithManager(ctx, mgr, auditor, restrictToNamespace); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "TrafficFilter")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "deployment.ec.kubeform.com",
		Version: "v1alpha1",
		Kind:    "TrafficFilterAssociation",
	}:
		if err := (&controllersdeployment.TrafficFilterAssociationReconciler{
			Client:   mgr.GetClient(),
			Log:      ctrl.Log.WithName("controllers").WithName("TrafficFilterAssociation"),
			Scheme:   mgr.GetScheme(),
			Gvk:      gvk,
			Provider: _provider,
			Resource: _provider.ResourcesMap["ec_deployment_traffic_filter_association"],
			TypeName: "ec_deployment_traffic_filter_association",
		}).SetupWithManager(ctx, mgr, auditor, restrictToNamespace); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "TrafficFilterAssociation")
			return err
		}

	default:
		return fmt.Errorf("Invalid CRD")
	}

	return nil
}

func SetupWebhook(mgr manager.Manager, gvk schema.GroupVersionKind) error {
	switch gvk {
	case schema.GroupVersionKind{
		Group:   "deployment.ec.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Deployment",
	}:
		if err := (&deploymentv1alpha1.Deployment{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "Deployment")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "deployment.ec.kubeform.com",
		Version: "v1alpha1",
		Kind:    "ElasticsearchKeystore",
	}:
		if err := (&deploymentv1alpha1.ElasticsearchKeystore{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "ElasticsearchKeystore")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "deployment.ec.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Extension",
	}:
		if err := (&deploymentv1alpha1.Extension{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "Extension")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "deployment.ec.kubeform.com",
		Version: "v1alpha1",
		Kind:    "TrafficFilter",
	}:
		if err := (&deploymentv1alpha1.TrafficFilter{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "TrafficFilter")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "deployment.ec.kubeform.com",
		Version: "v1alpha1",
		Kind:    "TrafficFilterAssociation",
	}:
		if err := (&deploymentv1alpha1.TrafficFilterAssociation{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "TrafficFilterAssociation")
			return err
		}

	default:
		return fmt.Errorf("Invalid Webhook")
	}

	return nil
}
