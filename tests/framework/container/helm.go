package container

import (
	"time"

	"helm.sh/helm/v3/pkg/action"

	fwclient "github.com/rancher/k3k/tests/framework/client"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// HelmInstaller provides configuration for Helm chart installation.
type HelmInstaller struct {
	ChartPath       string
	Namespace       string
	ReleaseName     string
	Timeout         time.Duration
	Wait            bool
	ControllerImage string
	KubeletImage    string
	KubeconfigPath  string
}

// InstallK3kChart installs the k3k Helm chart with the specified configuration.
// It uses the provided RESTClientGetter for authentication and returns the Helm action configuration.
func (h *HelmInstaller) InstallK3kChart(restClientGetter *fwclient.RESTClientGetter) *action.Configuration {
	_ = "STUB: not implemented"

	// Load chart
	return nil
}

// Initialize Helm action configuration

// Create install action

// Configure controller image

// Configure agent image

// Install chart

// NewHelmInstaller creates a new HelmInstaller with default values.
func NewHelmInstaller(controllerImage, kubeletImage, kubeconfigPath string) *HelmInstaller {
	_ = "STUB: not implemented"
	return nil
}
