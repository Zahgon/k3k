package policy

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	corev1 "k8s.io/api/core/v1"

	"github.com/rancher/k3k/pkg/apis/k3k.io/v1beta1"
)

const (
	PolicyNameLabelKey          = "policy.k3k.io/policy-name"
	ManagedByLabelKey           = "app.kubernetes.io/managed-by"
	VirtualPolicyControllerName = "k3k-policy-controller"

	policyFinalizerName = "policy.k3k.io/finalizer"
)

type VirtualClusterPolicyReconciler struct {
	Client      client.Client
	Scheme      *runtime.Scheme
	ClusterCIDR string
}

// Add the controller to manage the Virtual Cluster policies
func Add(mgr manager.Manager, clusterCIDR string, maxConcurrentReconciles int) error {
	_ = "STUB: not implemented"
	return nil
}

// namespaceEventHandler will enqueue a reconciliation of VCP when a Namespace changes
func namespaceEventHandler() handler.Funcs {
	_ = "STUB: not implemented"
	return *

	// When a Namespace is created, if it has the "policy.k3k.io/policy-name" label
	new(handler.Funcs)
}

// When a Namespace is updated, if it has the "policy.k3k.io/policy-name" label

// If labels haven't changed we can skip the reconciliation

// If No VCP before and after we can skip the reconciliation

// The VCP has not changed, but we enqueue a reconciliation because the PSA or other labels have changed

// Enqueue the old VCP name for cleanup

// Enqueue the new VCP name

// When a namespace is deleted all the resources in the namespace are deleted
// but we trigger the reconciliation to eventually perform some cluster-wide cleanup if necessary

// nodeEventHandler will enqueue a reconciliation of all the VCPs when a Node changes.
// This happens only if the ClusterCIDR is NOT specified, to handle the PodCIDRs in the NetworkPolicies.
func nodeEventHandler(r *VirtualClusterPolicyReconciler) handler.Funcs {
	_ = "STUB: not implemented"
	// enqueue all the available VirtualClusterPolicies
	return *new(handler.Funcs)
}

// Check if PodCIDR or PodCIDRs fields have changed.

// clusterEventHandler will enqueue a reconciliation of the VCP associated to the Namespace when a Cluster changes.
func clusterEventHandler(r *VirtualClusterPolicyReconciler) handler.Funcs {
	_ = "STUB: not implemented"
	return *new(handler.Funcs)
}

// When a Cluster is created, if its Namespace has the "policy.k3k.io/policy-name" label

// When a Cluster is updated, if its Namespace has the "policy.k3k.io/policy-name" label
// and if some of its spec influenced by the policy changed

// When a Cluster is deleted -> nothing to do

func (c *VirtualClusterPolicyReconciler) Reconcile(ctx context.Context, req reconcile.Request) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

// if DeletionTimestamp is not Zero -> finalize the object

// update Status if needed

// if there was an error during the reconciliation, return

// update VirtualClusterPolicy if needed

func (c *VirtualClusterPolicyReconciler) reconcileVirtualClusterPolicy(ctx context.Context, policy *v1beta1.VirtualClusterPolicy) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *VirtualClusterPolicyReconciler) reconcileMatchingNamespaces(ctx context.Context, policy *v1beta1.VirtualClusterPolicy) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *VirtualClusterPolicyReconciler) reconcileQuota(ctx context.Context, namespace string, policy *v1beta1.VirtualClusterPolicy) error {
	_ = "STUB: not implemented"
	return nil
}

// check if resourceQuota object exists and deletes it.

// create/update resource Quota

func (c *VirtualClusterPolicyReconciler) reconcileLimit(ctx context.Context, namespace string, policy *v1beta1.VirtualClusterPolicy) error {
	_ = "STUB: not implemented"
	return nil
}

// delete limitrange if spec.limits isnt specified.

func (c *VirtualClusterPolicyReconciler) reconcileClusters(ctx context.Context, namespace *corev1.Namespace, policy *v1beta1.VirtualClusterPolicy) error {
	_ = "STUB: not implemented"
	return nil
}

// continue updating also the other clusters even if an error occurred
