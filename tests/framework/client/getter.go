package client

import (
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// RESTClientGetter is a Kubernetes REST client getter implementation that satisfies
// the genericclioptions.RESTClientGetter interface. This is used primarily for Helm
// operations in tests.
type RESTClientGetter struct {
	clientconfig    clientcmd.ClientConfig
	restConfig      *rest.Config
	discoveryClient discovery.CachedDiscoveryInterface
}

// NewRESTClientGetter creates a new RESTClientGetter from kubeconfig bytes.
// This is used for Helm operations in tests.
func NewRESTClientGetter(kubeconfig []byte) (*RESTClientGetter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ToRESTConfig returns the REST config.
func (r *RESTClientGetter) ToRESTConfig() (*rest.Config, error) {
	_ = "STUB: not implemented"
	return nil,

		// ToDiscoveryClient returns the cached discovery client.
		nil
}

func (r *RESTClientGetter) ToDiscoveryClient() (discovery.CachedDiscoveryInterface, error) {
	_ = "STUB: not implemented"
	return *new(discovery.CachedDiscoveryInterface), nil
}

// ToRESTMapper returns a REST mapper from the discovery client.
func (r *RESTClientGetter) ToRESTMapper() (meta.RESTMapper, error) {
	_ = "STUB: not implemented"
	return *new(meta.RESTMapper), nil
}

// ToRawKubeConfigLoader returns the raw kubeconfig loader.
func (r *RESTClientGetter) ToRawKubeConfigLoader() clientcmd.ClientConfig {
	_ = "STUB: not implemented"
	return *new(clientcmd.ClientConfig)
}
