package cluster

import (
	"context"
	"crypto/tls"

	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	clientv3 "go.etcd.io/etcd/client/v3"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	ctrlruntimeclient "sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/rancher/k3k/pkg/apis/k3k.io/v1beta1"
)

const (
	statefulsetController = "k3k-statefulset-controller"
	etcdPodFinalizerName  = "etcdpod.k3k.io/finalizer"
)

type StatefulSetReconciler struct {
	Client ctrlruntimeclient.Client
	Scheme *runtime.Scheme
}

// Add adds a new controller to the manager
func AddStatefulSetController(ctx context.Context, mgr manager.Manager, maxConcurrentReconciles int) error {
	_ = "STUB: not implemented"
	// initialize a new Reconciler
	return nil
}

func (p *StatefulSetReconciler) Reconcile(ctx context.Context, req reconcile.Request) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

// we can ignore the IsNotFound error
// if the stateful set was deleted we have already cleaned up the pods

// If the StatefulSet is being deleted, we need to remove the finalizers from its pods
// and remove the finalizer from the StatefulSet itself.

// get cluster name from the object

func (p *StatefulSetReconciler) handleServerPod(ctx context.Context, cluster v1beta1.Cluster, pod *corev1.Pod) error {
	_ = "STUB: not implemented"
	return nil
}

// if etcd pod is marked for deletion then we need to remove it from the etcd member list before deletion

// check if cluster is deleted then remove the finalizer from the pod

// remove server from etcd

// remove our finalizer from the list and update it.

func (p *StatefulSetReconciler) getETCDTLS(ctx context.Context, cluster *v1beta1.Cluster) (*tls.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// create rootCA CertPool

// removePeer removes a peer from the cluster. The peer name and IP address must both match.
func removePeer(ctx context.Context, client *clientv3.Client, name, address string) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *StatefulSetReconciler) clusterToken(ctx context.Context, cluster *v1beta1.Cluster) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (p *StatefulSetReconciler) handleDeletion(ctx context.Context, sts *appsv1.StatefulSet) (ctrl.Result, error) {
	_ = "STUB: not implemented"
	return *new(ctrl.Result), nil
}

func (p *StatefulSetReconciler) listPods(ctx context.Context, sts *appsv1.StatefulSet) (*corev1.PodList, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
