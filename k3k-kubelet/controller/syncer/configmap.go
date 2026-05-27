package syncer

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	corev1 "k8s.io/api/core/v1"
)

const (
	configMapControllerName = "configmap-syncer"
	configMapFinalizerName  = "configmap.k3k.io/finalizer"
)

type ConfigMapSyncer struct {
	// SyncerContext contains all client information for host and virtual cluster
	*SyncerContext
}

func (c *ConfigMapSyncer) Name() string { _ = "STUB: not implemented"; return "" }

// AddConfigMapSyncer adds configmap syncer controller to the manager of the virtual cluster
func AddConfigMapSyncer(ctx context.Context, virtMgr, hostMgr manager.Manager, clusterName, clusterNamespace string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ConfigMapSyncer) filterResources(object client.Object) bool {
	_ = "STUB: not implemented"
	return false
}

// check for configMap Sync Config

// If syncing is disabled, only process deletions to allow for cleanup.

// Reconcile implements reconcile.Reconciler and synchronizes the objects in objs to the host cluster
func (c *ConfigMapSyncer) Reconcile(ctx context.Context, req reconcile.Request) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

// handle deletion

// deleting the synced configMap if exist

// remove the finalizer after cleaning up the synced configMap

// Add finalizer if it does not exist

// TODO: Add option to keep labels/annotation set by the host cluster

// translateConfigMap will translate a given configMap created in the virtual cluster and
// translates it to host cluster object
func (c *ConfigMapSyncer) translateConfigMap(configMap *corev1.ConfigMap) *corev1.ConfigMap {
	_ = "STUB: not implemented"
	return nil
}
