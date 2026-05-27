package cluster

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	ctrlruntimeclient "sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	podController = "k3k-pod-controller"
)

type PodReconciler struct {
	Client ctrlruntimeclient.Client
	Scheme *runtime.Scheme
}

// AddPodController adds a new controller for Pods to the manager.
// It will reconcile the Pods of the Host Cluster with the one of the Virtual Cluster.
func AddPodController(ctx context.Context, mgr manager.Manager, maxConcurrentReconciles int) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *PodReconciler) Reconcile(ctx context.Context, req reconcile.Request) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

// get cluster from the object
