package cmds

import (
	"context"
	"time"

	"github.com/spf13/cobra"
	"sigs.k8s.io/controller-runtime/pkg/client"

	corev1 "k8s.io/api/core/v1"

	"github.com/rancher/k3k/pkg/apis/k3k.io/v1beta1"
)

type CreateConfig struct {
	token                string
	clusterCIDR          string
	serviceCIDR          string
	servers              int
	agents               int
	serverArgs           []string
	agentArgs            []string
	serverEnvs           []string
	agentEnvs            []string
	labels               []string
	annotations          []string
	persistenceType      string
	storageClassName     string
	storageRequestSize   string
	version              string
	mode                 string
	kubeconfigServerHost string
	policy               string
	mirrorHostNodes      bool
	customCertsPath      string
	timeout              time.Duration
}

func NewClusterCreateCmd(appCtx *AppContext) *cobra.Command { _ = "STUB: not implemented"; return nil }

func createAction(appCtx *AppContext, config *CreateConfig) func(cmd *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

// add Host IP address as an extra TLS-SAN to expose the k3k cluster

// retry every 5s for at most 2m, or 25 times

func newCluster(name, namespace string, config *CreateConfig) (*v1beta1.Cluster, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func env(envSlice []string) []corev1.EnvVar { _ = "STUB: not implemented"; return nil }

func waitForClusterReconciled(ctx context.Context, k8sClient client.Client, cluster *v1beta1.Cluster, timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func waitForClusterReady(ctx context.Context, k8sClient client.Client, cluster *v1beta1.Cluster, timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// If resource ready -> stop polling

// If resource failed -> stop polling with an error

// Condition not met, continue polling.

func CreateCustomCertsSecrets(ctx context.Context, name, namespace, customCertsPath string, k8sclient client.Client) error {
	_ = "STUB: not implemented"
	return nil
}

func caCertSecret(certName, clusterName, clusterNamespace string, cert, key []byte) *corev1.Secret {
	_ = "STUB: not implemented"
	return nil
}

func parseKeyValuePairs(pairs []string, pairType string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

const clusterDetailsTemplate = `Cluster details:
  Mode: {{ .Mode }}
  Servers: {{ .Servers }}{{ if .Agents }}
  Agents: {{ .Agents }}{{ end }}
  Version: {{ if .Version }}{{ .Version }}{{ else }}{{ .HostVersion }}{{ end }} (Host: {{ .HostVersion }})
  Persistence:
    Type: {{.Persistence.Type}}{{ if .Persistence.StorageClassName }}
    StorageClass: {{ .Persistence.StorageClassName }}{{ end }}{{ if .Persistence.StorageRequestSize }}
    Size: {{ .Persistence.StorageRequestSize }}{{ end }}{{ if .Labels }}
  Labels: {{ range $key, $value := .Labels }}
    {{$key}}: {{$value}}{{ end }}{{ end }}{{ if .Annotations }}
  Annotations: {{ range $key, $value := .Annotations }}
    {{$key}}: {{$value}}{{ end }}{{ end }}`

func getClusterDetails(cluster *v1beta1.Cluster) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
