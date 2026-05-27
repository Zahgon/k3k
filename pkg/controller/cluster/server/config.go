package server

import (
	corev1 "k8s.io/api/core/v1"

	"github.com/rancher/k3k/pkg/apis/k3k.io/v1beta1"
)

// serverConfig are few options from k3s server options that will
// construct the yaml config file for k3s server
type serverConfig struct {
	ClusterCIDR        string   `yaml:"cluster-cidr,omitempty"`
	ClusterDNS         string   `yaml:"cluster-dns,omitempty"`
	ClusterInit        bool     `yaml:"cluster-init,omitempty"`
	DisableAgent       bool     `yaml:"disable-agent,omitempty"`
	Disable            []string `yaml:"disable,omitempty"`
	EgressSelectorMode string   `yaml:"egress-selector-mode,omitempty"`
	Server             string   `yaml:"server,omitempty"`
	ServiceCIDR        string   `yaml:"service-cidr,omitempty"`
	TLSSAN             []string `yaml:"tls-san,omitempty"`
	Token              string   `yaml:"token,omitempty"`
}

func (s *Server) Config(init bool, serviceIP string) (*corev1.Secret, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func buildServerConfig(cluster *v1beta1.Cluster, initServer bool, serviceIP, token string) serverConfig {
	_ = "STUB: not implemented"
	return *new(serverConfig)
}

func configSecretName(clusterName string, init bool) string { _ = "STUB: not implemented"; return "" }
