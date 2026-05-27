package server

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/client"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"

	"github.com/rancher/k3k/pkg/apis/k3k.io/v1beta1"
)

const (
	serverName     = "server"
	configName     = "server-config"
	initConfigName = "init-server-config"
)

// Server
type Server struct {
	cluster          *v1beta1.Cluster
	client           client.Client
	mode             string
	token            string
	image            string
	imagePullPolicy  string
	imagePullSecrets []string
}

func New(cluster *v1beta1.Cluster, client client.Client, token, image, imagePullPolicy string, imagePullSecrets []string) *Server {
	_ = "STUB: not implemented"
	return nil
}

func (s *Server) podSpec(ctx context.Context, image, name string, persistent bool, startupCmd string) corev1.PodSpec {
	_ = "STUB: not implemented"
	return *new(corev1.PodSpec)
}

// Use the server affinity from the policy status if it exists, otherwise fall back to the spec.

// Adding readiness probes to statefulset

// start the pod unprivileged in shared mode

// specify resource limits if specified for the servers.

// specifying ServerResources will take precedence over ServerLimits

// removing container previous limit

// add image pull secrets

func (s *Server) StatefulServer(ctx context.Context) (*appsv1.StatefulSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Server) setupDynamicPersistence() corev1.PersistentVolumeClaim {
	_ = "STUB: not implemented"
	return *new(corev1.PersistentVolumeClaim)
}

func (s *Server) setupStartCommand() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (s *Server) buildCABundleVolumes(ctx context.Context) ([]corev1.Volume, []corev1.VolumeMount, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Service account token secret is an exception (may not contain crt/key).

func (s *Server) mountCACert(volumeName, certName, secretName string, subPathMount string) (*corev1.Volume, []corev1.VolumeMount) {
	_ = "STUB: not implemented"
	return nil, nil
}

// avoid re-adding secretName in case of combined secret

// add the mount for the cert except for the service account token

// add the mount for the key

func (s *Server) buildAddonsVolumes(ctx context.Context) ([]corev1.Volume, []corev1.VolumeMount, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// skip creating the addon secret if it already exists and in the same namespace as the cluster

func sortedKeys(keyMap map[string]string) []string { _ = "STUB: not implemented"; return nil }
