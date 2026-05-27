package cmds

import (
	"github.com/spf13/cobra"

	"github.com/rancher/k3k/pkg/apis/k3k.io/v1beta1"
)

type UpdateConfig struct {
	servers     int32
	agents      int32
	labels      []string
	annotations []string
	version     string
	noConfirm   bool
}

func NewClusterUpdateCmd(appCtx *AppContext) *cobra.Command { _ = "STUB: not implemented"; return nil }

func updateFlags(cmd *cobra.Command, cfg *UpdateConfig) { _ = "STUB: not implemented"; return }

func updateAction(appCtx *AppContext, config *UpdateConfig) func(cmd *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

func confirmClusterUpdate(cluster *v1beta1.Cluster) bool { _ = "STUB: not implemented"; return false }
