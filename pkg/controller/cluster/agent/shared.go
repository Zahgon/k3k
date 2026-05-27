package agent

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	ctrlruntimeclient "sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/rancher/k3k/pkg/apis/k3k.io/v1beta1"
)

const (
	SharedNodeAgentName = "kubelet"
	SharedNodeMode      = "shared"
)

type SharedAgent struct {
	*Config
	serviceIP        string
	image            string
	imagePullPolicy  string
	imageRegistry    string
	token            string
	kubeletPort      int
	imagePullSecrets []string
}

type sharedAgentConfig struct {
	ClusterName      string `yaml:"clusterName"`
	ClusterNamespace string `yaml:"clusterNamespace"`
	KubeletPort      int    `yaml:"kubeletPort"`
	MirrorHostNodes  bool   `yaml:"mirrorHostNodes"`
	ServerIP         string `yaml:"serverIP"`
	ServiceName      string `yaml:"serviceName"`
	Token            string `yaml:"token"`
	Version          string `yaml:"version"`
}

func NewSharedAgent(config *Config, serviceIP, image, imagePullPolicy, token string, kubeletPort int, imagePullSecrets []string) *SharedAgent {
	_ = "STUB: not implemented"
	return nil
}

func (s *SharedAgent) Name() string { _ = "STUB: not implemented"; return "" }

func (s *SharedAgent) EnsureResources(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *SharedAgent) ensureObject(ctx context.Context, obj ctrlruntimeclient.Object) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *SharedAgent) config(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func sharedAgentData(cluster *v1beta1.Cluster, serviceName, token, ip string, kubeletPort int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *SharedAgent) daemonset(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (s *SharedAgent) podSpec(ctx context.Context) corev1.PodSpec {
	_ = "STUB: not implemented"
	return *new(corev1.PodSpec)
}

// Use the agent affinity from the policy status if it exists, otherwise fall back to the spec.

// specify resource limits if specified for the agents.

// specifying WorkerResources will take precedence over WorkerLimits

// removing container previous limit

func (s *SharedAgent) service(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (s *SharedAgent) dnsService(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (s *SharedAgent) serviceAccount(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *SharedAgent) role(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (s *SharedAgent) roleBinding(ctx context.Context) error { _ = "STUB: not implemented"; return nil }
