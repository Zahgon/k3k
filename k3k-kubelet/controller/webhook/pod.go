package webhook

import (
	"context"

	ctrlruntimeclient "sigs.k8s.io/controller-runtime/pkg/client"
)

const webhookName = "podmutating.k3k.io"

func RemovePodMutatingWebhook(ctx context.Context, virtualClient, hostClient ctrlruntimeclient.Client, clusterName, clusterNamespace string) error {
	_ = "STUB: not implemented"
	return nil
}
