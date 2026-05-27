package cmds

import (
	"context"

	"github.com/spf13/cobra"

	ctrlclient "sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/rancher/k3k/pkg/apis/k3k.io/v1beta1"
)

var keepData bool

func NewClusterDeleteCmd(appCtx *AppContext) *cobra.Command { _ = "STUB: not implemented"; return nil }

func delete(appCtx *AppContext) func(cmd *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

// keep bootstrap secrets and tokens if --keep-data flag is passed

// skip removing tokenSecret

func RemoveOwnerReferenceFromSecret(ctx context.Context, name string, cl ctrlclient.Client, cluster v1beta1.Cluster) error {
	_ = "STUB: not implemented"
	return nil
}
