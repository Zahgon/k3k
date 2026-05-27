package syncer

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	corev1 "k8s.io/api/core/v1"
)

const (
	secretControllerName = "secret-syncer"
	secretFinalizerName  = "secret.k3k.io/finalizer"
)

type SecretSyncer struct {
	// SyncerContext contains all client information for host and virtual cluster
	*SyncerContext
}

func (s *SecretSyncer) Name() string { _ = "STUB: not implemented"; return "" }

// AddSecretSyncer adds secret syncer controller to the manager of the virtual cluster
func AddSecretSyncer(ctx context.Context, virtMgr, hostMgr manager.Manager, clusterName, clusterNamespace string) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *SecretSyncer) filterResources(object client.Object) bool {
	_ = "STUB: not implemented"
	return false
}

// check for Secrets Sync Config

// If syncing is disabled, only process deletions to allow for cleanup.

// Reconcile implements reconcile.Reconciler and synchronizes the objects in objs to the host cluster
func (s *SecretSyncer) Reconcile(ctx context.Context, req reconcile.Request) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

// handle deletion

// deleting the synced secret if exist

// remove the finalizer after cleaning up the synced secret

// Add finalizer if it does not exist

// TODO: Add option to keep labels/annotation set by the host cluster

// translateSecret will translate a given secret created in the virtual cluster and
// translates it to host cluster object
func (s *SecretSyncer) translateSecret(secret *corev1.Secret) *corev1.Secret {
	_ = "STUB: not implemented"
	return nil
}
