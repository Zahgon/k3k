package main

import (
	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/rancher/k3k/pkg/log"
)

var (
	configFile string
	cfg        config
	logger     logr.Logger
	debug      bool
	logFormat  string
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "k3k-kubelet",
		Short: "virtual kubelet implementation k3k",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if err := InitializeConfig(cmd); err != nil {
				return err
			}

			logger = zapr.NewLogger(log.New(debug, logFormat))
			ctrl.SetLogger(logger)

			return nil
		},
		RunE: run,
	}

	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "", false, "Enable debug logging")
	rootCmd.PersistentFlags().StringVar(&logFormat, "log-format", "text", "Log format (text or json)")
	rootCmd.PersistentFlags().StringVar(&cfg.ClusterName, "cluster-name", "", "Name of the k3k cluster")
	rootCmd.PersistentFlags().StringVar(&cfg.ClusterNamespace, "cluster-namespace", "", "Namespace of the k3k cluster")
	rootCmd.PersistentFlags().StringVar(&cfg.Token, "token", "", "K3S token of the k3k cluster")
	rootCmd.PersistentFlags().StringVar(&cfg.HostKubeconfig, "host-kubeconfig", "", "Path to the host kubeconfig, if empty then virtual-kubelet will use incluster config")
	rootCmd.PersistentFlags().StringVar(&cfg.VirtKubeconfig, "virt-kubeconfig", "", "Path to the k3k cluster kubeconfig, if empty then virtual-kubelet will create its own config from k3k cluster")
	rootCmd.PersistentFlags().IntVar(&cfg.KubeletPort, "kubelet-port", 0, "kubelet API port number")
	rootCmd.PersistentFlags().StringVar(&cfg.ServiceName, "service-name", "", "The service name deployed by the k3k controller")
	rootCmd.PersistentFlags().StringVar(&cfg.AgentHostname, "agent-hostname", "", "Agent Hostname used for TLS SAN for the kubelet server")
	rootCmd.PersistentFlags().StringVar(&cfg.ServerIP, "server-ip", "", "Server IP used for registering the virtual kubelet to the cluster")
	rootCmd.PersistentFlags().StringVar(&cfg.Version, "version", "", "Version of kubernetes server")
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "/opt/rancher/k3k/config.yaml", "Path to k3k-kubelet config file")
	rootCmd.PersistentFlags().BoolVar(&cfg.MirrorHostNodes, "mirror-host-nodes", false, "Mirror real node objects from host cluster")

	if err := rootCmd.Execute(); err != nil {
		logrus.Fatal(err)
	}
}

func run(cmd *cobra.Command, args []string) error { _ = "STUB: not implemented"; return nil }

// InitializeConfig sets up viper to read from config file, environment variables, and flags.
// It uses a `flatcase` convention for viper keys to match the (lowercased) config file keys,
// while flags remain in kebab-case.
func InitializeConfig(cmd *cobra.Command) error {
	_ = "STUB: not implemented"

	// Bind every cobra flag to a viper key.
	// The viper key will be the flag name with dashes removed (flatcase).
	// e.g. "cluster-name" becomes "clustername"
	return nil
}

// Unmarshal all configuration into the global cfg struct.
// Viper correctly handles the precedence of flags > env > config.

// Separately get the debug flag, as it's not part of the main config struct.
