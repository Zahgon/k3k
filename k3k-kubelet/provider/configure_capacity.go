package provider

import (
	"context"
	"time"

	"github.com/go-logr/logr"
	"sigs.k8s.io/controller-runtime/pkg/client"

	corev1 "k8s.io/api/core/v1"

	"github.com/rancher/k3k/pkg/apis/k3k.io/v1beta1"
)

const (
	// UpdateNodeCapacityInterval is the interval at which the node capacity is updated.
	UpdateNodeCapacityInterval = 10 * time.Second
)

// milliScaleResources is a set of resource names that are measured in milli-units (e.g., CPU).
// This is used to determine whether to use MilliValue() for calculations.
var milliScaleResources = map[corev1.ResourceName]struct{}{
	corev1.ResourceCPU:                      {},
	corev1.ResourceMemory:                   {},
	corev1.ResourceStorage:                  {},
	corev1.ResourceEphemeralStorage:         {},
	corev1.ResourceRequestsCPU:              {},
	corev1.ResourceRequestsMemory:           {},
	corev1.ResourceRequestsStorage:          {},
	corev1.ResourceRequestsEphemeralStorage: {},
	corev1.ResourceLimitsCPU:                {},
	corev1.ResourceLimitsMemory:             {},
	corev1.ResourceLimitsEphemeralStorage:   {},
}

// coreResources is a set of the core infrastructure resource requests, if a quota resource
// is in this map then it should not be reflected to the virtual node capacity.
var coreResources = map[corev1.ResourceName]struct{}{
	corev1.ResourceCPU:                      {},
	corev1.ResourceMemory:                   {},
	corev1.ResourceStorage:                  {},
	corev1.ResourceEphemeralStorage:         {},
	corev1.ResourceRequestsCPU:              {},
	corev1.ResourceRequestsMemory:           {},
	corev1.ResourceRequestsStorage:          {},
	corev1.ResourceRequestsEphemeralStorage: {},
}

// StartNodeCapacityUpdater starts a goroutine that periodically updates the capacity
// of the virtual node based on host node capacity and any applied ResourceQuotas.
func startNodeCapacityUpdater(ctx context.Context, logger logr.Logger, hostClient client.Client, virtualClient client.Client, virtualCluster v1beta1.Cluster, virtualNodeName string) {
	_ = "STUB: not implemented"
	return
}

// updateNodeCapacity will update the virtual node capacity (and the allocatable field) with the sum of all the resource in the host nodes.
// If the nodeLabels are specified only the matching nodes will be considered.
func updateNodeCapacity(ctx context.Context, logger logr.Logger, hostClient client.Client, virtualClient client.Client, virtualCluster v1beta1.Cluster, virtualNodeName string) {
	_ = "STUB: not implemented"
	// by default we get the resources of the same Node where the kubelet is running
	return
}

// we need to check if the virtual cluster resources are "limited" through ResourceQuotas
// If so we will use the minimum resources

// get the node's quota and merge it with the current values

// mergeQuotas takes multiple resource quotas lists and returns a single list that represents
// the most restrictive set of resource quotas. For each resource name, it selects the minimum
// quantity found across all the provided lists.
func mergeQuotas(resourceLists ...corev1.ResourceList) corev1.ResourceList {
	_ = "STUB: not implemented"
	return *new(corev1.ResourceList)
}

// If it's the first time we see it OR the new one is smaller -> Update

// distributeQuotas divides the total resource quotas among all active virtual nodes,
// capped by each node's actual host capacity. This ensures that each virtual node
// reports a fair share of the available resources without exceeding what its
// underlying host node can provide.
//
// For each resource type the algorithm uses a multi-pass redistribution loop:
//  1. Divide the remaining quota evenly among eligible nodes (sorted by name for
//     determinism), assigning any integer remainder to the first nodes alphabetically.
//  2. Cap each node's share at its host allocatable capacity.
//  3. Remove nodes that have reached their host capacity.
//  4. If there is still unallocated quota (because some nodes were capped below their
//     even share), repeat from step 1 with the remaining quota and remaining nodes.
//
// The loop terminates when the quota is fully distributed or no eligible nodes remain.
func distributeQuotas(hostResourceMap, virtResourceMap map[string]corev1.ResourceList, quotas corev1.ResourceList) map[string]corev1.ResourceList {
	_ = "STUB: not implemented"
	return nil
}

// fill out any allocatable resource that does not exist in quota

// Distribute each resource type from the policy's hard quota

// eligible nodes for each distribution cycle

// Populate the host nodes capacity map and the initial effective nodes

// skip the node if the resource does not exist on the host node

// Start of the distribution cycle, each cycle will distribute the quota resource
// evenly between nodes, each node can not exceed the corresponding host node capacity

// We cap the quantity to the hostNode capacity

// filterQuotas filters a resource list from any resource that is not eligible to be used for node capacity
// like core resources requests, it also strips requests/limits prefixes from other extended resources
// for example "requests.nvidia.com/gpu" will return back "nvidia.com/gpu"
func filterQuotas(resources corev1.ResourceList) corev1.ResourceList {
	_ = "STUB: not implemented"
	return *new(corev1.ResourceList)
}
