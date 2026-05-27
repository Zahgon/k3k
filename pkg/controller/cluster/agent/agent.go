package agent

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"

	ctrlruntimeclient "sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/rancher/k3k/pkg/apis/k3k.io/v1beta1"
)

const (
	configName = "agent-config"
)

type ResourceEnsurer interface {
	EnsureResources(context.Context) error
}

type Config struct {
	cluster *v1beta1.Cluster
	client  ctrlruntimeclient.Client
	scheme  *runtime.Scheme
}

func NewConfig(cluster *v1beta1.Cluster, client ctrlruntimeclient.Client, scheme *runtime.Scheme) *Config {
	_ = "STUB: not implemented"
	return nil
}

func configSecretName(clusterName string) string { _ = "STUB: not implemented"; return "" }

func ensureObject(ctx context.Context, cfg *Config, obj ctrlruntimeclient.Object) error {
	_ = "STUB: not implemented"
	return nil
}
