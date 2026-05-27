package cluster

import (
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

func newClusterPredicate() predicate.Predicate {
	_ = "STUB: not implemented"
	return *new(predicate.Predicate)
}

func clusterNamespacedName(object client.Object) types.NamespacedName {
	_ = "STUB: not implemented"
	return *new(types.NamespacedName)
}
