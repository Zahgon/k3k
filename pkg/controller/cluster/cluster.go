package cluster

import (
	"context"
	"errors"
	"time"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/tools/record"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	corev1 "k8s.io/api/core/v1"

	"github.com/rancher/k3k/pkg/apis/k3k.io/v1beta1"
	"github.com/rancher/k3k/pkg/controller/cluster/agent"
	"github.com/rancher/k3k/pkg/controller/cluster/server"
)

const (
	clusterController    = "k3k-cluster-controller"
	clusterFinalizerName = "cluster.k3k.io/finalizer"
	ClusterInvalidName   = "system"

	SyncEnabledLabelKey = "k3k.io/sync-enabled"
	SyncSourceLabelKey  = "k3k.io/sync-source"
	SyncSourceHostLabel = "host"

	defaultVirtualClusterCIDR = "10.52.0.0/16"
	defaultVirtualServiceCIDR = "10.53.0.0/16"
	defaultSharedClusterCIDR  = "10.42.0.0/16"
	defaultSharedServiceCIDR  = "10.43.0.0/16"
	memberRemovalTimeout      = time.Minute * 1

	storageClassEnabledIndexField       = "spec.sync.storageClasses.enabled"
	storageClassStatusEnabledIndexField = "status.policy.sync.storageClasses.enabled"
)

var (
	ErrClusterValidation         = errors.New("cluster validation error")
	ErrCustomCACertSecretMissing = errors.New("custom CA certificate secret is missing")
)

type Config struct {
	ClusterCIDR                 string
	SharedAgentImage            string
	SharedAgentImagePullPolicy  string
	VirtualAgentImage           string
	VirtualAgentImagePullPolicy string
	K3SServerImage              string
	K3SServerImagePullPolicy    string
	ServerImagePullSecrets      []string
	AgentImagePullSecrets       []string
}

type ClusterReconciler struct {
	DiscoveryClient *discovery.DiscoveryClient
	Client          client.Client
	Scheme          *runtime.Scheme
	PortAllocator   *agent.PortAllocator

	record.EventRecorder
	Config
}

// Add adds a new controller to the manager
func Add(ctx context.Context, mgr manager.Manager, config *Config, maxConcurrentReconciles int, portAllocator *agent.PortAllocator, eventRecorder record.EventRecorder) error {
	_ = "STUB: not implemented"
	return nil
}

// initialize a new Reconciler

// index the 'spec.sync.storageClasses.enabled' field

// index the 'status.policy.sync.storageClasses.enabled' field

func (r *ClusterReconciler) mapStorageClassToCluster(ctx context.Context, obj client.Object) []reconcile.Request {
	_ = "STUB: not implemented"
	return nil
}

// Merge and deduplicate clusters

func namespaceEventHandler(r *ClusterReconciler) handler.Funcs {
	_ = "STUB: not implemented"
	return *

	// We don't need to update for create or delete events
	new(handler.Funcs)
}

// When a Namespace is updated, if it has the "policy.k3k.io/policy-name" label

// If policy hasn't changed we can skip the reconciliation

// Enqueue all the Cluster in the namespace

func (c *ClusterReconciler) Reconcile(ctx context.Context, req reconcile.Request) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

// if DeletionTimestamp is not Zero -> finalize the object

// Set initial status if not already set

// add finalizer

// if there was an error during the reconciliation, return

// update Cluster if needed

func (c *ClusterReconciler) reconcileCluster(ctx context.Context, cluster *v1beta1.Cluster) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ClusterReconciler) reconcile(ctx context.Context, cluster *v1beta1.Cluster) error {
	_ = "STUB: not implemented"
	return nil
}

// if the Version is not specified we will try to use the same Kubernetes version of the host.
// This version is stored in the Status object, and it will not be updated if already set.

// update Status HostVersion

// in shared mode try to lookup the serviceCIDR

// in virtual mode assign a default serviceCIDR

// Important: if you need to call the Server API of the Virtual Cluster
// this needs to be done AFTER he kubeconfig has been generated

// ensureBootstrapSecret will create or update the Secret containing the bootstrap data from the k3s server
func (c *ClusterReconciler) ensureBootstrapSecret(ctx context.Context, cluster *v1beta1.Cluster, serviceIP, token string) error {
	_ = "STUB: not implemented"
	return nil
}

// ensureKubeconfigSecret will create or update the Secret containing the kubeconfig data from the k3s server
func (c *ClusterReconciler) ensureKubeconfigSecret(ctx context.Context, cluster *v1beta1.Cluster, serviceIP string, port int) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ClusterReconciler) createClusterConfigs(ctx context.Context, cluster *v1beta1.Cluster, server *server.Server, serviceIP string) error {
	_ = "STUB: not implemented"
	// create init node config
	return nil
}

// create servers configuration

func (c *ClusterReconciler) ensureNetworkPolicy(ctx context.Context, cluster *v1beta1.Cluster) error {
	_ = "STUB: not implemented"
	return nil
}

// network policies are managed by the Policy -> delete the one created as a standalone cluster

func (c *ClusterReconciler) ensureClusterService(ctx context.Context, cluster *v1beta1.Cluster) (*corev1.Service, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *ClusterReconciler) ensureIngress(ctx context.Context, cluster *v1beta1.Cluster) error {
	_ = "STUB: not implemented"
	return nil
}

// delete existing Ingress if Expose or IngressConfig are nil

func (c *ClusterReconciler) ensureStorageClasses(ctx context.Context, cluster *v1beta1.Cluster) error {
	_ = "STUB: not implemented"
	return nil
}

// If a policy is applied to the virtual cluster we need to use its SyncConfig, if available

// If storageclass sync is disabled, clean up any managed storage classes.

// filter the StorageClasses disabled for the sync, and the one not matching the selector

// if sync is disabled -> continue

// if selector doesn't match -> continue
// an empty selector matche everything

// delete StorageClasses with the sync disabled

func (c *ClusterReconciler) server(ctx context.Context, cluster *v1beta1.Cluster, server *server.Server) error {
	_ = "STUB: not implemented"
	return nil

	// create headless service for the statefulset
}

// Add the finalizer to the StatefulSet so the statefulset controller can handle cleanup.

func (c *ClusterReconciler) bindClusterRoles(ctx context.Context, cluster *v1beta1.Cluster) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ClusterReconciler) ensureAgent(ctx context.Context, cluster *v1beta1.Cluster, serviceIP, token string) error {
	_ = "STUB: not implemented"
	return nil
}

// Assign port from pool if shared agent enabled mirroring of host nodes

func (c *ClusterReconciler) validate(cluster *v1beta1.Cluster, policy v1beta1.VirtualClusterPolicy) error {
	_ = "STUB: not implemented"
	return nil
}

// lookupServiceCIDR attempts to determine the cluster's service CIDR.
// It first attempts to create a failing Service (with an invalid cluster IP)and extracts the expected CIDR from the resulting error.
// If that fails, it searches the 'kube-apiserver' Pod's arguments for the --service-cluster-ip-range flag.
func (c *ClusterReconciler) lookupServiceCIDR(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "",

		// Try to look for the serviceCIDR creating a failing service.
		// The error should contain the expected serviceCIDR
		nil
}

// validate serviceCIDR

// Try to look for the the kube-apiserver Pod, and look for the '--service-cluster-ip-range' flag.

// validate serviceCIDR

// validateCustomCACerts will make sure that all the cert secrets exists
func (c *ClusterReconciler) validateCustomCACerts(credentialSources v1beta1.CredentialSources) error {
	_ = "STUB: not implemented"
	return nil
}
