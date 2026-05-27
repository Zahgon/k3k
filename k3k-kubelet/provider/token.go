package provider

import (
	"context"

	authv1 "k8s.io/api/authentication/v1"
	corev1 "k8s.io/api/core/v1"
)

const (
	kubeAPIAccessPrefix          = "kube-api-access"
	serviceAccountTokenMountPath = "/var/run/secrets/kubernetes.io/serviceaccount"
)

// transformTokens copies the serviceaccount tokens used by virtualPod's serviceaccount to a secret on the host cluster and mount it
// to look like the serviceaccount token
func (p *Provider) transformTokens(ctx context.Context, virtualPod, hostPod *corev1.Pod) error {
	_ = "STUB: not implemented"
	return nil
}

// transform projected service account token

// transform kube-api-access token for all containers in virtualPod

func (p *Provider) transformKubeAccessToken(ctx context.Context, virtualPod, hostPod *corev1.Pod) error {
	_ = "STUB: not implemented"
	// skip this process if the kube-api-access is already removed from the pod
	// this is needed in case users already adds their own custom tokens like in rancher imported clusters
	return nil
}

// extracting the tokens data from the secret we just created

// To avoid race conditions we need to check if the secret's data has been populated
// including the token, ca.crt and namespace

// transformProjectedTokens will iterate over the host pod projected volume sources
// and transform projected tokens to use a requested token secret from the virtual cluster
// instead the automatically generated secret on the host cluster.
func (p *Provider) transformProjectedTokens(ctx context.Context, virtualPod, hostPod *corev1.Pod) error {
	_ = "STUB: not implemented"
	return nil
}

// replace the projected token volume with a projected secret

func (p *Provider) requestTokenSecret(ctx context.Context, token *corev1.ServiceAccountTokenProjection, virtualPod *corev1.Pod) (*corev1.Secret, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// create a virtual secret with that token

// creating unique name for the virtual secret based on the request attributes

func virtualSecret(name, namespace, serviceAccountName string) *corev1.Secret {
	_ = "STUB: not implemented"
	return nil
}

func (p *Provider) translateAndCreateHostTokenSecret(ctx context.Context, projectedToken *corev1.Secret) (*corev1.Secret, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func hasKubeAccessVolume(pod *corev1.Pod) bool { _ = "STUB: not implemented"; return false }

func removeKubeAccessVolume(pod *corev1.Pod) { _ = "STUB: not implemented"; return }

// init containers

// ephemeral containers

func addKubeAccessVolume(pod *corev1.Pod, hostSecretName string) { _ = "STUB: not implemented"; return }

func generateTokenSecretName(serviceAccountName, tokenPath string, tokenReq *authv1.TokenRequest) string {
	_ = "STUB: not implemented"
	return ""
}
