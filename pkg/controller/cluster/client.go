package cluster

import (
	"context"

	ctrlruntimeclient "sigs.k8s.io/controller-runtime/pkg/client"
)

// newVirtualClient creates a new Client that can be used to interact with the virtual cluster
func newVirtualClient(ctx context.Context, hostClient ctrlruntimeclient.Client, clusterName, clusterNamespace string) (ctrlruntimeclient.Client, error) {
	_ = "STUB: not implemented"
	return *new(ctrlruntimeclient.Client), nil
}
