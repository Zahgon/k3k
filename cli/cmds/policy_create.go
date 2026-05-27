package cmds

import (
	"context"

	"github.com/spf13/cobra"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/rancher/k3k/pkg/apis/k3k.io/v1beta1"
)

type VirtualClusterPolicyCreateConfig struct {
	mode        string
	labels      []string
	annotations []string
	namespaces  []string
	overwrite   bool
}

func NewPolicyCreateCmd(appCtx *AppContext) *cobra.Command { _ = "STUB: not implemented"; return nil }

func policyCreateAction(appCtx *AppContext, config *VirtualClusterPolicyCreateConfig) func(cmd *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

func createNamespace(ctx context.Context, client client.Client, name, policyName string) error {
	_ = "STUB: not implemented"
	return nil
}

func createPolicy(ctx context.Context, client client.Client, config *VirtualClusterPolicyCreateConfig, policyName string) (*v1beta1.VirtualClusterPolicy, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func bindPolicyToNamespaces(ctx context.Context, client client.Client, config *VirtualClusterPolicyCreateConfig, policyName string) error {
	_ = "STUB: not implemented"
	return nil
}

// same policy found, no need to update

// no old policy, safe to update

// different policy, warn or check for overwrite flag
