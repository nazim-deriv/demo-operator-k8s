/*
Copyright 2023.

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

package controllers

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	testoperatorv1alpha1 "github.com/nazim-deriv/demo-operator-k8s/api/v1alpha1"
)

// ResourceOperatorReconciler reconciles a ResourceOperator object
type ResourceOperatorReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=test-operator.deriv.services,resources=resourceoperators,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=test-operator.deriv.services,resources=resourceoperators/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=test-operator.deriv.services,resources=resourceoperators/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the ResourceOperator object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.1/pkg/reconcile
func (r *ResourceOperatorReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)

	ResourceOperator := &testoperatorv1alpha1.ResourceOperator{}
	if err := r.Get(ctx, req.NamespacedName, ResourceOperator); err != nil {
		log.Error(err, "Failed to fetch ResourceOperator", "Name", req.Name, "Namespace", req.Namespace)
		return reconcile.Result{}, err
	}

	log.Info("ResourceOperator created", "Name", ResourceOperator.Name)

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ResourceOperatorReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&testoperatorv1alpha1.ResourceOperator{}).
		Complete(r)
}
