package client

import (
	"context"

	"github.com/testcontainers/testcontainers-go/modules/k3s"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// Config holds the Kubernetes client configuration and clients.
type Config struct {
	RestConfig *rest.Config
	Clientset  *kubernetes.Clientset
	Client     client.Client
	HostIP     string
}

// InitFromKubeconfig initializes Kubernetes clients from the KUBECONFIG environment variable.
// It sets up logging, reads the kubeconfig file, creates REST config and clients.
// The scheme parameter should be created using the scheme package.
func InitFromKubeconfig(ctx context.Context, scheme *runtime.Scheme, k3sContainer *k3s.K3sContainer) (*Config, error) {
	_ = "STUB: not implemented"
	// Setup logger
	return nil, nil
}

// Get kubeconfig path from environment

// Read kubeconfig file

// InitFromBytes initializes Kubernetes clients from kubeconfig bytes.
// The scheme parameter should be created using the scheme package.
func InitFromBytes(ctx context.Context, kubeconfig []byte, scheme *runtime.Scheme, k3sContainer *k3s.K3sContainer) (*Config, error) {
	_ = "STUB: not implemented"
	// Create REST config from kubeconfig
	return nil, nil
}

// Extract host IP from REST config

// Create Kubernetes clientset

// Create controller-runtime client

// getServerIP extracts the server IP from the REST config.
// If running with testcontainers, it returns the container IP.
// Otherwise, it parses the hostname from the REST config host.
func getServerIP(ctx context.Context, cfg *rest.Config, k3sContainer *k3s.K3sContainer) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// If Host includes a port, u.Hostname() extracts just the hostname part
