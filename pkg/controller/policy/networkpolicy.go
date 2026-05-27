package policy

import (
	"context"

	networkingv1 "k8s.io/api/networking/v1"

	"github.com/rancher/k3k/pkg/apis/k3k.io/v1beta1"
)

func (c *VirtualClusterPolicyReconciler) reconcileNetworkPolicy(ctx context.Context, namespace string, policy *v1beta1.VirtualClusterPolicy) error {
	_ = "STUB: not implemented"
	return nil
}

// if disabled then delete the existing network policy

// otherwise try to create/update

func networkPolicy(namespaceName string, policy *v1beta1.VirtualClusterPolicy, cidrList []string) *networkingv1.NetworkPolicy {
	_ = "STUB: not implemented"
	return nil
}
