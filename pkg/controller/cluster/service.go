package cluster

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	ctrlruntimeclient "sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	serviceController = "k3k-service-controller"
)

type ServiceReconciler struct {
	HostClient ctrlruntimeclient.Client
}

// Add adds a new controller to the manager
func AddServiceController(ctx context.Context, mgr manager.Manager, maxConcurrentReconciles int) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *ServiceReconciler) Reconcile(ctx context.Context, req reconcile.Request) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

// Some services are owned by the cluster but don't have the annotations set (i.e. the kubelet svc)
// They don't exists in the virtual cluster, so we can skip them

// get cluster from the object
