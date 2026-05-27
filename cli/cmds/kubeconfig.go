package cmds

import (
	"github.com/spf13/cobra"

	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"

	"github.com/rancher/k3k/pkg/apis/k3k.io/v1beta1"
)

type GenerateKubeconfigConfig struct {
	name                 string
	configName           string
	cn                   string
	org                  []string
	altNames             []string
	expirationDays       int64
	kubeconfigServerHost string
}

func NewKubeconfigCmd(appCtx *AppContext) *cobra.Command { _ = "STUB: not implemented"; return nil }

func NewKubeconfigGenerateCmd(appCtx *AppContext) *cobra.Command {
	_ = "STUB: not implemented"
	return nil
}

func generateKubeconfigFlags(cmd *cobra.Command, cfg *GenerateKubeconfigConfig) {
	_ = "STUB: not implemented"
	return
}

func generate(appCtx *AppContext, cfg *GenerateKubeconfigConfig) func(cmd *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

func writeKubeconfigFile(cluster *v1beta1.Cluster, kubeconfig *clientcmdapi.Config, configName string) error {
	_ = "STUB: not implemented"
	return nil
}
