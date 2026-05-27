package server

import (
	"context"

	networkingv1 "k8s.io/api/networking/v1"

	"github.com/rancher/k3k/pkg/apis/k3k.io/v1beta1"
)

const (
	httpsPort     = 443
	k3sServerPort = 6443
	etcdPort      = 2379
)

func IngressName(clusterName string) string { _ = "STUB: not implemented"; return "" }

func Ingress(ctx context.Context, cluster *v1beta1.Cluster) networkingv1.Ingress {
	_ = "STUB: not implemented"
	return *new(networkingv1.Ingress)
}

func ingressRules(cluster *v1beta1.Cluster) []networkingv1.IngressRule {
	_ = "STUB: not implemented"
	return nil
}
