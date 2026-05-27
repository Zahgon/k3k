package kubeconfig

import (
	"context"
	"time"

	"sigs.k8s.io/controller-runtime/pkg/client"

	certutil "github.com/rancher/dynamiclistener/cert"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"

	"github.com/rancher/k3k/pkg/apis/k3k.io/v1beta1"
)

type KubeConfig struct {
	AltNames   certutil.AltNames
	CN         string
	ORG        []string
	ExpiryDate time.Duration
}

func New() *KubeConfig { _ = "STUB: not implemented"; return nil }

func (k *KubeConfig) Generate(ctx context.Context, client client.Client, cluster *v1beta1.Cluster, hostServerIP string, port int) (*clientcmdapi.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewConfig(url string, serverCA, clientCert, clientKey []byte) *clientcmdapi.Config {
	_ = "STUB: not implemented"
	return nil
}

func getURLFromService(ctx context.Context, client client.Client, cluster *v1beta1.Cluster, hostServerIP string, serverPort int) (string, error) {
	_ = "STUB: not implemented"
	// get the server service to extract the right IP
	return "", nil
}

// if ingress is specified, use the ingress host
