package networkaddonsconfig

import (
	"log"
	"time"

	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

var resyncPeriod = 5 * time.Minute

// newPodReconciler returns a new reconcile.Reconciler
func newPodReconciler() *ReconcilePods {
	return &ReconcilePods{}
}

var _ reconcile.Reconciler = &ReconcilePods{}

// ReconcilePods watches for updates to specified resources and then updates its StatusManager
type ReconcilePods struct {
	resources []types.NamespacedName
}

func (r *ReconcilePods) SetResources(resources []types.NamespacedName) {
	r.resources = resources
}

// Reconcile updates the NetworkAddonsConfig.Status to match the current state of the
// watched Deployments/DaemonSets
func (r *ReconcilePods) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	found := false
	for _, name := range r.resources {
		if name.Namespace == request.Namespace && name.Name == request.Name {
			found = true
			break
		}
	}
	if !found {
		return reconcile.Result{}, nil
	}

	log.Printf("Reconciling update to %s/%s\n", request.Namespace, request.Name)

	return reconcile.Result{RequeueAfter: resyncPeriod}, nil
}
