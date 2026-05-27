package policy

import (
	"context"

	corev1 "k8s.io/api/core/v1"

	"github.com/rancher/k3k/pkg/apis/k3k.io/v1beta1"
)

// reconcileNamespacePodSecurityLabels will update the labels of the namespace to reconcile the PSA level specified in the VirtualClusterPolicy
func (c *VirtualClusterPolicyReconciler) reconcileNamespacePodSecurityLabels(ctx context.Context, namespace *corev1.Namespace, policy *v1beta1.VirtualClusterPolicy) {
	_ = "STUB: not implemented"
	return
}

// cleanup of old labels

// if a PSA level is specified add the proper labels

// skip the 'warn' only for the privileged PSA level

// cleanupNamespaces will cleanup the Namespaces without the "policy.k3k.io/policy-name" label
// deleting the resources in them with the "app.kubernetes.io/managed-by=k3k-policy-controller" label
func (c *VirtualClusterPolicyReconciler) cleanupNamespaces(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// This will match all the resources managed by the K3k Policy controller
// that have the app.kubernetes.io/managed-by=k3k-policy-controller label

// If the namespace is not bound to any policy, or if the policy it was bound to no longer exists,
// we need to clear policy-related fields on its Cluster objects.

// if the namespace is bound to a policy -> cleanup resources of other policies

// log the error but continue cleaning up the other namespaces

// clearPolicyFieldsForClustersInNamespace sets the policy status on Cluster objects in the given namespace to nil.
func (c *VirtualClusterPolicyReconciler) clearPolicyFieldsForClustersInNamespace(ctx context.Context, namespace string) error {
	_ = "STUB: not implemented"
	return nil
}
