package cluster

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	"github.com/rancher/k3k/pkg/apis/k3k.io/v1beta1"
)

func (c *ClusterReconciler) finalizeCluster(ctx context.Context, cluster *v1beta1.Cluster) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

// Set the Terminating phase and condition

// Deallocate ports for kubelet if used

// delete API server lease

// Remove finalizer from the cluster and update it only when all resources are cleaned up

func (c *ClusterReconciler) unbindClusterRoles(ctx context.Context, cluster *v1beta1.Cluster) error {
	_ = "STUB: not implemented"
	return nil
}

// remove the clusterSubject from the ClusterRoleBinding
