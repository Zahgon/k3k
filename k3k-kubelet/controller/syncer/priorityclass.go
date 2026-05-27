package syncer

import (
	"context"
	"strings"

	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	schedulingv1 "k8s.io/api/scheduling/v1"
	ctrlruntimeclient "sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	PriorityClassGlobalDefaultAnnotation = "priorityclass.k3k.io/globalDefault"

	priorityClassControllerName = "priorityclass-syncer-controller"
	priorityClassFinalizerName  = "priorityclass.k3k.io/finalizer"
)

type PriorityClassSyncer struct {
	*SyncerContext
}

// AddPriorityClassSyncer adds a PriorityClass reconciler to k3k-kubelet
func AddPriorityClassSyncer(ctx context.Context, virtMgr, hostMgr manager.Manager, clusterName, clusterNamespace string) error {
	_ = "STUB: not implemented"
	// initialize a new Reconciler
	return nil
}

// IgnoreSystemPrefixPredicate filters out resources whose names start with "system-".
var ignoreSystemPrefixPredicate = predicate.Funcs{
	UpdateFunc: func(e event.UpdateEvent) bool {
		return !strings.HasPrefix(e.ObjectOld.GetName(), "system-")
	},
	CreateFunc: func(e event.CreateEvent) bool {
		return !strings.HasPrefix(e.Object.GetName(), "system-")
	},
	DeleteFunc: func(e event.DeleteEvent) bool {
		return !strings.HasPrefix(e.Object.GetName(), "system-")
	},
	GenericFunc: func(e event.GenericEvent) bool {
		return !strings.HasPrefix(e.Object.GetName(), "system-")
	},
}

func (r *PriorityClassSyncer) filterResources(object ctrlruntimeclient.Object) bool {
	_ = "STUB: not implemented"
	return false
}

// check for priorityClassConfig

// If syncing is disabled, only process deletions to allow for cleanup.

func (r *PriorityClassSyncer) Reconcile(ctx context.Context, req reconcile.Request) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

// handle deletion

// deleting the synced service if exists
// TODO add test for previous implementation without err != nil check, and also check the other controllers

// remove the finalizer after cleaning up the synced service

// Add finalizer if it does not exist

// create the priorityClass on the host

func (r *PriorityClassSyncer) translatePriorityClass(priorityClass schedulingv1.PriorityClass) *schedulingv1.PriorityClass {
	_ = "STUB: not implemented"
	return nil
}
