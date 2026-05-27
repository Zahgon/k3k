package agent

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	ctrlruntimeclient "sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	VirtualNodeMode      = "virtual"
	virtualNodeAgentName = "agent"
)

type VirtualAgent struct {
	*Config
	serviceIP        string
	token            string
	Image            string
	ImagePullPolicy  string
	ImageRegistry    string
	imagePullSecrets []string
}

type virtualAgentConfig struct {
	Server     string `yaml:"server"`
	Token      string `yaml:"token"`
	WithNodeId bool   `yaml:"with-node-id"`
}

func NewVirtualAgent(config *Config, serviceIP, token, Image, ImagePullPolicy string, imagePullSecrets []string) *VirtualAgent {
	_ = "STUB: not implemented"
	return nil
}

func (v *VirtualAgent) Name() string { _ = "STUB: not implemented"; return "" }

func (v *VirtualAgent) EnsureResources(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *VirtualAgent) ensureObject(ctx context.Context, obj ctrlruntimeclient.Object) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *VirtualAgent) config(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func virtualAgentData(serviceIP, token string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (v *VirtualAgent) deployment(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (v *VirtualAgent) podSpec(ctx context.Context, image, name string) corev1.PodSpec {
	_ = "STUB: not implemented"
	return *new(corev1.PodSpec)
}

// Use the agent affinity from the policy status if it exists, otherwise fall back to the spec.

// specify resource limits if specified for the servers.

// specifying WorkerResources will take precedence over WorkerLimits

// removing container previous limit
