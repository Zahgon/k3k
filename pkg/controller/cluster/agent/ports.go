package agent

import (
	"context"

	"k8s.io/kubernetes/pkg/apis/core"
	"k8s.io/kubernetes/pkg/registry/core/service/portallocator"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	corev1 "k8s.io/api/core/v1"
	ctrlruntimeclient "sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	kubeletPortRangeConfigMapName = "k3k-kubelet-port-range"

	rangeKey          = "range"
	allocatedPortsKey = "allocatedPorts"
	snapshotDataKey   = "snapshotData"
)

type PortAllocator struct {
	ctrlruntimeclient.Client

	KubeletCM *corev1.ConfigMap
}

func NewPortAllocator(ctx context.Context, client ctrlruntimeclient.Client) (*PortAllocator, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *PortAllocator) InitPortAllocatorConfig(ctx context.Context, client ctrlruntimeclient.Client, kubeletPortRange string) manager.Runnable {
	_ = "STUB: not implemented"
	return *new(manager.Runnable)
}

func (a *PortAllocator) getOrCreate(ctx context.Context, configmap *corev1.ConfigMap, portRange string) error {
	_ = "STUB: not implemented"
	return nil
}

// creating the configMap for the first time

func (a *PortAllocator) AllocateKubeletPort(ctx context.Context, clusterName, clusterNamespace string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (a *PortAllocator) DeallocateKubeletPort(ctx context.Context, clusterName, clusterNamespace string, kubeletPort int) error {
	_ = "STUB: not implemented"
	return nil
}

// allocatePort will assign port to the cluster from a port Range configured for k3k
func (a *PortAllocator) allocatePort(ctx context.Context, clusterName, clusterNamespace string, configMap *corev1.ConfigMap) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// get configMap first to avoid conflicts

// allocate a new port and save the snapshot

// deallocatePort will remove the port used by the cluster from the port range
func (a *PortAllocator) deallocatePort(ctx context.Context, clusterName, clusterNamespace string, configMap *corev1.ConfigMap, port int) error {
	_ = "STUB: not implemented"
	return nil
}

// check if the cluster already exists in the configMap

// parsePortMap will convert ConfigMap Data to a portMap of string keys and values of ints
func parsePortMap(portMapData string) (map[string]int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// serializePortMap will convert a portMap of string keys and values of ints to ConfigMap Data
func serializePortMap(m map[string]int) (string, error) { _ = "STUB: not implemented"; return "", nil }

func saveSnapshot(portAllocator *portallocator.PortAllocator, snapshot *core.RangeAllocation, configMap *corev1.ConfigMap, portsMap map[string]int) error {
	_ = "STUB: not implemented"
	// save the new snapshot
	return nil
}

// update the configmap with the new portsMap and the new snapshot
