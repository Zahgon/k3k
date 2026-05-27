package cluster

import (
	"context"

	"github.com/rancher/k3k/pkg/apis/k3k.io/v1beta1"
)

const (
	// Condition Types
	ConditionReady = "Ready"

	// Condition Reasons
	ReasonValidationFailed   = "ValidationFailed"
	ReasonProvisioning       = "Provisioning"
	ReasonProvisioned        = "Provisioned"
	ReasonProvisioningFailed = "ProvisioningFailed"
	ReasonTerminating        = "Terminating"
)

func (c *ClusterReconciler) updateStatus(ctx context.Context, cluster *v1beta1.Cluster, reconcileErr error) {
	_ = "STUB: not implemented"
	return
}

// Handle validation errors specifically to set the Pending phase.

// If there's an error, but it's not a validation error, the cluster is in a failed state.

// If we reach here, everything is successful.

// Only emit event on transition to Ready
