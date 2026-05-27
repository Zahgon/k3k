package container

import (
	"context"

	"github.com/testcontainers/testcontainers-go/modules/k3s"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// SetupK3s creates and starts a K3s testcontainer with the specified version and loads the provided images.
// It returns the container instance and a temporary kubeconfig file path.
// The kubeconfig is automatically cleaned up via DeferCleanup.
func SetupK3s(ctx context.Context, k3sVersion, controllerImage, kubeletImage string) (*k3s.K3sContainer, string) {
	_ = "STUB: not implemented"
	return nil, ""
}

// Normalize version string (replace + with -)

// Start K3s container

// Get kubeconfig from container

// Write kubeconfig to temp file

// Load images into the K3s container

// Register cleanup

// Set KUBECONFIG environment variable
