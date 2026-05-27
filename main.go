//go:generate ./scripts/generate
package main

import (
	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/runtime"

	clientgoscheme "k8s.io/client-go/kubernetes/scheme"

	"github.com/rancher/k3k/cli/cmds"
	"github.com/rancher/k3k/pkg/apis/k3k.io/v1beta1"
	"github.com/rancher/k3k/pkg/buildinfo"
	"github.com/rancher/k3k/pkg/controller/cluster"
	"github.com/rancher/k3k/pkg/log"
)

var (
	scheme                  = runtime.NewScheme()
	config                  cluster.Config
	kubeconfig              string
	kubeletPortRange        string
	maxConcurrentReconciles int
	debug                   bool
	logFormat               string
	logger                  logr.Logger
)

func init() {
	_ = clientgoscheme.AddToScheme(scheme)
	_ = v1beta1.AddToScheme(scheme)
}

func main() {
	rootCmd := &cobra.Command{
		Use:     "k3k",
		Short:   "k3k controller",
		Version: buildinfo.Version,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return validate()
		},
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			cmds.InitializeConfig(cmd)

			logger = zapr.NewLogger(log.New(debug, logFormat))
		},
		RunE: run,
	}

	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "", false, "Debug level logging")
	rootCmd.PersistentFlags().StringVar(&logFormat, "log-format", "text", "Log format (text or json)")
	rootCmd.PersistentFlags().StringVar(&kubeconfig, "kubeconfig", "", "kubeconfig path")
	rootCmd.PersistentFlags().StringVar(&config.ClusterCIDR, "cluster-cidr", "", "Cluster CIDR to be added to the networkpolicy")
	rootCmd.PersistentFlags().StringVar(&config.SharedAgentImage, "agent-shared-image", "rancher/k3k-kubelet", "K3K Virtual Kubelet image")
	rootCmd.PersistentFlags().StringVar(&config.SharedAgentImagePullPolicy, "agent-shared-image-pull-policy", "", "K3K Virtual Kubelet image pull policy must be one of Always, IfNotPresent or Never")
	rootCmd.PersistentFlags().StringVar(&config.VirtualAgentImage, "agent-virtual-image", "rancher/k3s", "K3S Virtual Agent image")
	rootCmd.PersistentFlags().StringVar(&config.VirtualAgentImagePullPolicy, "agent-virtual-image-pull-policy", "", "K3S Virtual Agent image pull policy must be one of Always, IfNotPresent or Never")
	rootCmd.PersistentFlags().StringVar(&kubeletPortRange, "kubelet-port-range", "50000-51000", "Port Range for k3k kubelet in shared mode")
	rootCmd.PersistentFlags().StringVar(&config.K3SServerImage, "k3s-server-image", "rancher/k3s", "K3K server image")
	rootCmd.PersistentFlags().StringVar(&config.K3SServerImagePullPolicy, "k3s-server-image-pull-policy", "", "K3K server image pull policy")
	rootCmd.PersistentFlags().StringSliceVar(&config.ServerImagePullSecrets, "server-image-pull-secret", nil, "Image pull secret used for for servers")
	rootCmd.PersistentFlags().StringSliceVar(&config.AgentImagePullSecrets, "agent-image-pull-secret", nil, "Image pull secret used for for agents")
	rootCmd.PersistentFlags().IntVar(&maxConcurrentReconciles, "max-concurrent-reconciles", 50, "maximum number of concurrent reconciles")

	if err := rootCmd.Execute(); err != nil {
		logger.Error(err, "failed to run k3k controller")
	}
}

func run(cmd *cobra.Command, args []string) error { _ = "STUB: not implemented"; return nil }

func validate() error { _ = "STUB: not implemented"; return nil }
