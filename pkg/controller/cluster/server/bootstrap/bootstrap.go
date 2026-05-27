package bootstrap

import (
	"context"
	"errors"

	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/rancher/k3k/pkg/apis/k3k.io/v1beta1"
)

var ErrServerNotReady = errors.New("server not ready")

type ControlRuntimeBootstrap struct {
	ServerCA        content `json:"serverCA"`
	ServerCAKey     content `json:"serverCAKey"`
	ClientCA        content `json:"clientCA"`
	ClientCAKey     content `json:"clientCAKey"`
	ETCDServerCA    content `json:"etcdServerCA"`
	ETCDServerCAKey content `json:"etcdServerCAKey"`
}

type content struct {
	Timestamp string
	Content   string
}

// Generate generates the bootstrap for the cluster:
// 1- use the server token to get the bootstrap data from k3s
// 2- save the bootstrap data as a secret
func GenerateBootstrapData(ctx context.Context, cluster *v1beta1.Cluster, ip, token string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func requestBootstrap(token, serverIP string) (*ControlRuntimeBootstrap, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func basicAuth(username, password string) string { _ = "STUB: not implemented"; return "" }

func decodeBootstrap(bootstrap *ControlRuntimeBootstrap) error {
	_ = "STUB: not implemented"
	// client-ca
	return nil
}

// client-ca-key

// server-ca

// server-ca-key

// etcd-ca

// etcd-ca-key

func DecodedBootstrap(token, ip string) (*ControlRuntimeBootstrap, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetFromSecret(ctx context.Context, client client.Client, cluster *v1beta1.Cluster) (*ControlRuntimeBootstrap, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
