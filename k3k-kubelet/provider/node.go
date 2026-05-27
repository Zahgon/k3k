package provider

import (
	"context"

	corev1 "k8s.io/api/core/v1"
)

// Node implements the node.Provider interface from Virtual Kubelet
type Node struct {
	notifyCallback func(*corev1.Node)
}

// Ping is called to check if the node is healthy - in the current format it always is
func (n *Node) Ping(context.Context) error {
	_ = "STUB: not implemented"

	// NotifyNodeStatus sets the callback function for a node being changed. As of now, no changes are made
	return nil
}

func (n *Node) NotifyNodeStatus(ctx context.Context, cb func(*corev1.Node)) {
	_ = "STUB: not implemented"
	return
}
