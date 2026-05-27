package syncer

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	corev1 "k8s.io/api/core/v1"
	ctrlruntimeclient "sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	pvcControllerName = "pvc-syncer-controller"
	pvcFinalizerName  = "pvc.k3k.io/finalizer"
	pseudoPVLabel     = "pod.k3k.io/pseudoPV"
)

type PVCReconciler struct {
	*SyncerContext
}

// AddPVCSyncer adds persistentvolumeclaims syncer controller to k3k-kubelet
func AddPVCSyncer(ctx context.Context, virtMgr, hostMgr manager.Manager, clusterName, clusterNamespace string) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *PVCReconciler) filterResources(object ctrlruntimeclient.Object) bool {
	_ = "STUB: not implemented"
	return false
}

// check for pvc config

// If syncing is disabled, only process deletions to allow for cleanup.

func (r *PVCReconciler) Reconcile(ctx context.Context, req reconcile.Request) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

// handle deletion

// deleting the synced pvc if exists

// delete the synced virtual PV

// remove the finalizer after cleaning up the synced pvc

// Add finalizer if it does not exist

// create the pvc on host

// note that we dont need to update the PVC on the host cluster, only syncing the PVC to allow being
// handled by the host cluster.

// Creating a virtual PV to bound the existing PVC in the virtual cluster - needed for scheduling of
// the consumer pods

func (r *PVCReconciler) pvc(obj *corev1.PersistentVolumeClaim) *corev1.PersistentVolumeClaim {
	_ = "STUB: not implemented"
	return nil
}

func (r *PVCReconciler) createVirtualPersistentVolume(ctx context.Context, pvc corev1.PersistentVolumeClaim) error {
	_ = "STUB: not implemented"
	return nil
}

// Update spec to set the volumeName binding

func newPersistentVolume(obj *corev1.PersistentVolumeClaim) *corev1.PersistentVolume {
	_ = "STUB: not implemented"
	return nil
}
