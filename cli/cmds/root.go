package cmds

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type AppContext struct {
	RestConfig *rest.Config
	Client     client.Client

	// Global flags
	Debug      bool
	Kubeconfig string
	namespace  string
}

func NewRootCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func (ctx *AppContext) Namespace(name string) string { _ = "STUB: not implemented"; return "" }

func loadRESTConfig(kubeconfig string) (*rest.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func CobraFlagNamespace(appCtx *AppContext, flag *pflag.FlagSet) { _ = "STUB: not implemented"; return }

func InitializeConfig(cmd *cobra.Command) { _ = "STUB: not implemented"; return }

// Bind the current command's flags to viper

// Apply the viper config value to the flag when the flag is not set and viper has a value
