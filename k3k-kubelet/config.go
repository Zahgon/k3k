package main

// config has all virtual-kubelet startup options
type config struct {
	ClusterName      string `mapstructure:"clusterName"`
	ClusterNamespace string `mapstructure:"clusterNamespace"`
	ServiceName      string `mapstructure:"serviceName"`
	Token            string `mapstructure:"token"`
	AgentHostname    string `mapstructure:"agentHostname"`
	HostKubeconfig   string `mapstructure:"hostKubeconfig"`
	VirtKubeconfig   string `mapstructure:"virtKubeconfig"`
	KubeletPort      int    `mapstructure:"kubeletPort"`
	ServerIP         string `mapstructure:"serverIP"`
	Version          string `mapstructure:"version"`
	MirrorHostNodes  bool   `mapstructure:"mirrorHostNodes"`
}

func (c *config) validate() error { _ = "STUB: not implemented"; return nil }
