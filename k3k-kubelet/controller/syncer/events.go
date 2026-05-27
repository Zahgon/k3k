package syncer

import (
	"context"

	"k8s.io/client-go/tools/record"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

const (
	eventControllerName = "event-syncer"
)

type EventSyncer struct {
	// virtEventRecorder is a K8s EventRecorder to emit events into the
	// virtual cluster.
	virtEventRecorder record.EventRecorder

	// SyncerContext contains all client information for host and virtual
	// cluster.
	*SyncerContext
}

func (s *EventSyncer) Name() string { _ = "STUB: not implemented"; return "" }

// AddEventSyncer adds event syncer controller to the manager of the virtual
// cluster.
func AddEventSyncer(ctx context.Context, virtMgr, hostMgr manager.Manager, clusterName, clusterNamespace string, virtEventRecorder record.EventRecorder) error {
	_ = "STUB: not implemented"
	return nil
}

// Reconcile implements reconcile.Reconciler and synchronizes relevant events
// between the host and virtual clusters.
func (s *EventSyncer) Reconcile(ctx context.Context, req reconcile.Request) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

// Look up the corresponding object in the virtual cluster.

// translateEventMessage replaces occurrences of the host-cluster pod name and
// "namespace/name" pattern in the event message with the corresponding
// virtual-cluster values, so the emitted event refers to virtual coordinates.
func translateEventMessage(message, hostName, hostNamespace, virtualName, virtualNamespace string) string {
	_ = "STUB: not implemented"
	// Replace the combined "namespace/name" first to avoid a partial match
	// turning the namespace portion into the virtual namespace before the name
	// is replaced.
	return ""
}

// Replace any remaining standalone occurrences of the host pod name.
