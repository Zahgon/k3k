package syncer

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	corev1 "k8s.io/api/core/v1"
	ctrlruntimeclient "sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	serviceControllerName = "service-syncer-controller"
	serviceFinalizerName  = "service.k3k.io/finalizer"
)

type ServiceReconciler struct {
	*SyncerContext
}

// AddServiceSyncer adds service syncer controller to the manager of the virtual cluster
func AddServiceSyncer(ctx context.Context, virtMgr, hostMgr manager.Manager, clusterName, clusterNamespace string) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *ServiceReconciler) Reconcile(ctx context.Context, req reconcile.Request) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

// handle deletion

// deleting the synced service if exists

// remove the finalizer after cleaning up the synced service

// Add finalizer if it does not exist

// create or update the service on host

func (r *ServiceReconciler) filterResources(object ctrlruntimeclient.Object) bool {
	_ = "STUB: not implemented"
	return false
}

// check for serviceSyncConfig

// If syncing is disabled, only process deletions to allow for cleanup.

func (s *ServiceReconciler) service(obj *corev1.Service) *corev1.Service {
	_ = "STUB: not implemented"
	return nil
}

// don't sync finalizers to the host
