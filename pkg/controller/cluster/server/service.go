package server

import (
	corev1 "k8s.io/api/core/v1"

	"github.com/rancher/k3k/pkg/apis/k3k.io/v1beta1"
)

func Service(cluster *v1beta1.Cluster) *corev1.Service { _ = "STUB: not implemented"; return nil }

// If no expose is specified, default to ClusterIP

// If expose is specified, set the type to the appropriate type

// default to clusterIP for ingress or empty expose config

// addLoadBalancerPorts adds the load balancer ports to the service
func addLoadBalancerPorts(service *corev1.Service, loadbalancerConfig v1beta1.LoadBalancerConfig, k3sServerPort, etcdPort corev1.ServicePort) {
	_ = "STUB: not implemented"
	// If the server port is not specified, use the default port
	return
}

// If the server port is specified, set the port, otherwise the service will not be exposed

// If the etcd port is not specified, use the default port

// If the etcd port is specified, set the port, otherwise the service will not be exposed

// addNodePortPorts adds the node port ports to the service
func addNodePortPorts(service *corev1.Service, nodePortConfig v1beta1.NodePortConfig, k3sServerPort, etcdPort corev1.ServicePort) {
	_ = "STUB: not implemented"
	// If the server port is not specified Kubernetes will set the node port to a random port between 30000-32767
	return
}

// If the server port is in the range of 30000-32767, set the node port
// otherwise the service will not be exposed

// If the etcd port is not specified Kubernetes will set the node port to a random port between 30000-32767

// If the etcd port is in the range of 30000-32767, set the node port
// otherwise the service will not be exposed

func (s *Server) StatefulServerService() *corev1.Service { _ = "STUB: not implemented"; return nil }

func ServiceName(clusterName string) string { _ = "STUB: not implemented"; return "" }

func headlessServiceName(clusterName string) string { _ = "STUB: not implemented"; return "" }
