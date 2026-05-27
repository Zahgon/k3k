package cluster

import (
	"context"

	corev1 "k8s.io/api/core/v1"

	"github.com/rancher/k3k/pkg/apis/k3k.io/v1beta1"
)

func (c *ClusterReconciler) token(ctx context.Context, cluster *v1beta1.Cluster) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// get token data from secretRef

func (c *ClusterReconciler) ensureTokenSecret(ctx context.Context, cluster *v1beta1.Cluster) (string, error) {
	_ = "STUB: not implemented"
	return "",

		// check if the secret is already created
		nil
}

func random(size int) (string, error) { _ = "STUB: not implemented"; return "", nil }

func TokenSecretObj(token, name, namespace string) corev1.Secret {
	_ = "STUB: not implemented"
	return *new(corev1.Secret)
}

func TokenSecretName(clusterName string) string { _ = "STUB: not implemented"; return "" }
