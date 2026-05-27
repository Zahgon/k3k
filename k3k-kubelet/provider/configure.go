package provider

import (
	"github.com/go-logr/logr"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	corev1 "k8s.io/api/core/v1"

	"github.com/rancher/k3k/pkg/apis/k3k.io/v1beta1"
)

func ConfigureNode(logger logr.Logger, node *corev1.Node, hostname string, servicePort int, ip string, hostMgr manager.Manager, virtualClient client.Client, virtualCluster v1beta1.Cluster, version string, mirrorHostNodes bool) error {
	_ = "STUB: not implemented"
	return nil
}

// configure versions

// nodeConditions returns the basic conditions which mark the node as ready
func nodeConditions() []corev1.NodeCondition { _ = "STUB: not implemented"; return nil }
