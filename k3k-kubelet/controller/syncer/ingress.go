package syncer

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	networkingv1 "k8s.io/api/networking/v1"
	ctrlruntimeclient "sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	ingressControllerName = "ingress-syncer-controller"
	ingressFinalizerName  = "ingress.k3k.io/finalizer"
)

type IngressReconciler struct {
	*SyncerContext
}

// AddIngressSyncer adds ingress syncer controller to the manager of the virtual cluster
func AddIngressSyncer(ctx context.Context, virtMgr, hostMgr manager.Manager, clusterName, clusterNamespace string) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *IngressReconciler) filterResources(object ctrlruntimeclient.Object) bool {
	_ = "STUB: not implemented"
	return false
}

// check for ingressConfig

// If syncing is disabled, only process deletions to allow for cleanup.

func (r *IngressReconciler) Reconcile(ctx context.Context, req reconcile.Request) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

// If a policy is applied to the virtual cluster we need to use its SyncConfig, if available

// handle deletion

// deleting the synced service if exists

// remove the finalizer after cleaning up the synced service

// Add finalizer if it does not exist

// create or update the ingress on host

func (s *IngressReconciler) ingress(obj *networkingv1.Ingress, disableTLSSecretTranslation bool) *networkingv1.Ingress {
	_ = "STUB: not implemented"
	return nil
}

// modify services in rules to point to the synced services

// TLS Secret translation disable, return early without translating TLS secrets in the ingress spec

// ensure tls secrets are also translated
