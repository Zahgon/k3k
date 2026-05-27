package provider

import (
	"context"
	"errors"
	"io"

	"github.com/go-logr/logr"
	"github.com/virtual-kubelet/virtual-kubelet/node/api"
	"github.com/virtual-kubelet/virtual-kubelet/node/nodeutil"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	dto "github.com/prometheus/client_model/go"
	corev1 "k8s.io/api/core/v1"
	cv1 "k8s.io/client-go/kubernetes/typed/core/v1"
	v1alpha1stats "k8s.io/kubelet/pkg/apis/stats/v1alpha1"

	"github.com/rancher/k3k/k3k-kubelet/translate"
)

// check at compile time if the Provider implements the nodeutil.Provider interface
var _ nodeutil.Provider = (*Provider)(nil)

// ClusterContext includes the controller runtime manager and clients
type ClusterContext struct {
	Config     rest.Config
	Client     client.Client
	CoreClient cv1.CoreV1Interface
	Manager    manager.Manager
}

// Provider implements nodetuil.Provider from virtual Kubelet.
// TODO: Implement NotifyPods and the required usage so that this can be an async provider
type Provider struct {
	Host             ClusterContext
	Virtual          ClusterContext
	Translator       translate.ToHostTranslator
	ClusterNamespace string
	ClusterName      string
	serverIP         string
	dnsIP            string
	agentHostname    string
	logger           logr.Logger
}

var ErrRetryTimeout = errors.New("provider timed out")

func New(hostConfig rest.Config, hostMgr, virtualMgr manager.Manager, logger logr.Logger, namespace, name, serverIP, dnsIP, agentHostname string) (*Provider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetContainerLogs retrieves the logs of a container by name from the provider.
func (p *Provider) GetContainerLogs(ctx context.Context, namespace, name, containerName string, opts api.ContainerLogOpts) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

// RunInContainer executes a command in a container in the pod, copying data
// between in/out/err and the container's stdin/stdout/stderr.
func (p *Provider) RunInContainer(ctx context.Context, namespace, name, containerName string, cmd []string, attach api.AttachIO) error {
	_ = "STUB: not implemented"
	return nil
}

// AttachToContainer attaches to the executing process of a container in the pod, copying data
// between in/out/err and the container's stdin/stdout/stderr.
func (p *Provider) AttachToContainer(ctx context.Context, namespace, name, containerName string, attach api.AttachIO) error {
	_ = "STUB: not implemented"
	return nil
}

// GetStatsSummary gets the stats for the node, including running pods
func (p *Provider) GetStatsSummary(ctx context.Context) (*v1alpha1stats.Summary, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// skip pods that are not in the cluster namespace

// rewrite the PodReference to match the data of the virtual cluster

// GetMetricsResource gets the metrics for the node, including running pods
func (p *Provider) GetMetricsResource(ctx context.Context) ([]*dto.MetricFamily, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PortForward forwards a local port to a port on the pod
func (p *Provider) PortForward(ctx context.Context, namespace, name string, port int32, stream io.ReadWriteCloser) error {
	_ = "STUB: not implemented"
	return nil
}

// Today this doesn't work properly. When the port ward is supposed to stop, the caller (this provider)
// should send a value on stopChannel so that the PortForward is stopped. However, we only have a ReadWriteCloser
// so more work is needed to detect a close and handle that appropriately.

// CreatePod executes createPod with retry
func (p *Provider) CreatePod(ctx context.Context, pod *corev1.Pod) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Provider) createPod(ctx context.Context, pod *corev1.Pod) error {
	_ = "STUB: not implemented"
	return nil
}

// get Cluster definition

// get Pod from Virtual Cluster

// Copy the virtual Pod and use it as a baseline for the hostPod
// do some basic translation and clearing some values (UID, ResourceVersion, ...)

// Clear the NodeName to allow scheduling, and set affinity to prefer scheduling the Pod on the same host node as the virtual kubelet,
// unless the user has specified their own affinity, in which case the user's affinity is respected.

// The pod's own nodeSelector is ignored.
// The final selector is determined by the cluster spec, but overridden by a policy if present.

// setting the hostname for the pod if its not set

// When a PriorityClass is set we will use the translated one in the HostCluster.
// If the Cluster or a Policy defines a PriorityClass of the host we are going to use that one.
// Note: the core-dns and local-path-provisioner pod are scheduled by k3s with the
// 'system-cluster-critical' and 'system-node-critical' default priority classes.
//
// TODO: we probably need to define a custom "intermediate" k3k-system-* priority

// if the priority class is set we need to remove the priority

// volumes will often refer to resources in the virtual cluster
// but instead need to refer to the synced host cluster version

// sync serviceaccount token to a the host cluster

// inject networking information to the pod including the virtual cluster controlplane endpoint

// set ownerReference to the cluster object

// withRetry retries passed function with interval and timeout
func (p *Provider) withRetry(ctx context.Context, f func(context.Context, *corev1.Pod) error, pod *corev1.Pod) error {
	_ = "STUB: not implemented"
	return nil
}

// retryFn will retry until the operation succeed, or the timeout occurs

// log that the retry failed?

// transformVolumes changes the volumes to the representation in the host cluster
func (p *Provider) transformVolumes(podNamespace string, volumes []corev1.Volume) {
	_ = "STUB: not implemented"
	return
}

// Skip volumes related to Kube API access

// UpdatePod executes updatePod with retry
func (p *Provider) UpdatePod(ctx context.Context, pod *corev1.Pod) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Provider) updatePod(ctx context.Context, pod *corev1.Pod) error {
	_ = "STUB: not implemented"
	// Once scheduled a Pod cannot update other fields than the image of the containers, initcontainers and a few others
	// See: https://kubernetes.io/docs/concepts/workloads/pods/#pod-update-and-replacement
	return nil
}

//
//	Host Pod update
//

// Ephemeral containers update (subresource)

//
//	Virtual Pod update
//

// Ephemeral containers update (subresource)

func updatePod(dst, src *corev1.Pod) { _ = "STUB: not implemented"; return }

// updateContainerImages will update the images of the original container images with the same name
func updateContainerImages(dst, src []corev1.Container) { _ = "STUB: not implemented"; return }

// DeletePod executes deletePod with retry
func (p *Provider) DeletePod(ctx context.Context, pod *corev1.Pod) error {
	_ = "STUB: not implemented"
	return nil
}

// deletePod takes a Kubernetes Pod and deletes it from the provider. Once a pod is deleted, the provider is
// expected to call the NotifyPods callback with a terminal pod status where all the containers are in a terminal
// state, as well as the pod. DeletePod may be called multiple times for the same pod.
func (p *Provider) deletePod(ctx context.Context, pod *corev1.Pod) error {
	_ = "STUB: not implemented"
	return nil
}

// GetPod retrieves a pod by name from the provider (can be cached).
// The Pod returned is expected to be immutable, and may be accessed
// concurrently outside of the calling goroutine. Therefore it is recommended
// to return a version after DeepCopy.
func (p *Provider) GetPod(ctx context.Context, namespace, name string) (*corev1.Pod, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetPodStatus retrieves the status of a pod by name from the provider.
// The PodStatus returned is expected to be immutable, and may be accessed
// concurrently outside of the calling goroutine. Therefore it is recommended
// to return a version after DeepCopy.
func (p *Provider) GetPodStatus(ctx context.Context, namespace, name string) (*corev1.PodStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Provider) getPodFromHostCluster(ctx context.Context, hostPodName string) (*corev1.Pod, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetPods retrieves a list of all pods running on the provider (can be cached).
// The Pods returned are expected to be immutable, and may be accessed
// concurrently outside of the calling goroutine. Therefore it is recommended
// to return a version after DeepCopy.
func (p *Provider) GetPods(ctx context.Context) ([]*corev1.Pod, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// configureNetworking will inject network information to each pod to connect them to the
// virtual cluster api server, as well as confiugre DNS information to connect them to the
// synced coredns on the host cluster.
func configureNetworking(pod *corev1.Pod, podName, podNamespace, serverIP, dnsIP string) {
	_ = "STUB: not implemented"
	// inject serverIP to hostalias for the pod
	return
}

// injecting cluster DNS IP to the pods except for coredns pod

// inject networking information to the pod's environment variables

// handle init containers as well

// handle ephemeral containers as well

// mergeEnvVars will override the orig environment variables if found in the updated list and will add them to the list if not found
func mergeEnvVars(orig, updated []corev1.EnvVar) []corev1.EnvVar {
	_ = "STUB: not implemented"
	return nil
}

// create map for single lookup

// Remove the updated variable from the map

// Any variables remaining in the map are new and should be appended to the original slice.

func (p *Provider) configurePodEnvs(hostPod, virtualPod *corev1.Pod) {
	_ = "STUB: not implemented"
	return
}

func (p *Provider) configureEnv(virtualPod *corev1.Pod, envs []corev1.EnvVar) []corev1.EnvVar {
	_ = "STUB: not implemented"
	return nil
}

// for name and namespace we need to hardcode the virtual cluster values, and clear the FieldRef

func (p *Provider) configureEnvFrom(virtualPod *corev1.Pod, envs []corev1.EnvFromSource) []corev1.EnvFromSource {
	_ = "STUB: not implemented"
	return nil
}
