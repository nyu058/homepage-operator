/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"context"

	"k8s.io/apimachinery/pkg/api/errors"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	homepagev1 "nathanyu.com/homepage-operator/api/v1"
)

// HomePageEntryReconciler reconciles a HomePageEntry object
type HomePageEntryReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=homepage.nathanyu.com,resources=homepageentries,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=homepage.nathanyu.com,resources=homepageentries/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=homepage.nathanyu.com,resources=homepageentries/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the HomePageEntry object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.17.3/pkg/reconcile
func (r *HomePageEntryReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// Fetch the HomePageEntry resource
	entry := &homepagev1.HomePageEntry{}
	err := r.Get(context.Background(), req.NamespacedName, entry)
	if err != nil {
		// Error handling
	}

	// Get Services and Ingresses referenced in the HomePageEntry
	services := &corev1.ServiceList{}
	err = r.List(context.Background(), services)
	if err != nil {
		// Error handling
	}

	ingresses := &networkingv1.IngressList{}
	err = r.List(context.Background(), ingresses)
	if err != nil {
		// Error handling
	}

	// Create HTML content with links to Services and Ingresses
	htmlContent := "<html><body>"
	for _, service := range services.Items {
		htmlContent += "<a href='" + service.Spec.ClusterIP + "'>" + service.Name + "</a><br>"
	}
	for _, ingress := range ingresses.Items {
		for _, rule := range ingress.Spec.Rules {
			htmlContent += "<a href='" + rule.Host + "'>" + ingress.Name + "</a><br>"
		}
	}
	htmlContent += "</body></html>"

	// Create or update a ConfigMap with the HTML content
	configMap := &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "homepage",
			Namespace: req.Namespace,
		},
		Data: map[string]string{
			"index.html": htmlContent,
		},
	}
	existingConfigMap := &corev1.ConfigMap{}
	err = r.Get(context.Background(), client.ObjectKey{Namespace: configMap.Namespace, Name: configMap.Name}, existingConfigMap)
	if err != nil {
		if errors.IsNotFound(err) {
			// ConfigMap doesn't exist, create a new one
			err = r.Create(context.Background(), configMap)
			if err != nil {
				// Error handling
			}
		} else {
			// Error handling
		}
	} else {
		// ConfigMap exists, update it
		existingConfigMap.Data = configMap.Data
		err = r.Update(context.Background(), existingConfigMap)
		if err != nil {
			// Error handling
		}
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *HomePageEntryReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&homepagev1.HomePageEntry{}).
		Complete(r)
}
