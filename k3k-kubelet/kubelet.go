package main

import (
	"context"

	"github.com/go-logr/logr"
	"github.com/virtual-kubelet/virtual-kubelet/node/nodeutil"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/record"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	ctrlruntimeclient "sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/rancher/k3k/pkg/apis/k3k.io/v1beta1"
)

var baseScheme = runtime.NewScheme()

func init() {
	_ = clientgoscheme.AddToScheme(baseScheme)
	_ = v1beta1.AddToScheme(baseScheme)
}

type kubelet struct {
	virtualCluster v1beta1.Cluster

	name       string
	port       int
	hostConfig *rest.Config
	virtConfig *rest.Config
	agentIP    string
	dnsIP      string
	hostClient ctrlruntimeclient.Client
	virtClient kubernetes.Interface
	hostMgr    manager.Manager
	virtualMgr manager.Manager
	node       *nodeutil.Node
	logger     logr.Logger
	token      string

	virtEventRecorder record.EventRecorder
	eb                record.EventBroadcaster
}

func newKubelet(ctx context.Context, c *config) (*kubelet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// virtual client will only use core types (for now), no need to add anything other than the basics

// get the cluster's DNS IP to be injected to pods

func clusterIP(ctx context.Context, serviceName, clusterNamespace string, hostClient ctrlruntimeclient.Client) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (k *kubelet) start(ctx context.Context) {
	_ = "STUB: not implemented"
	// any one of the following 3 tasks (host manager, virtual manager, node) crashing will stop the
	// program, and all 3 of them block on start, so we start them here in go-routines
	return
}

// run the node async so that we can wait for it to be ready in another call

func (k *kubelet) newProviderFunc(cfg config) nodeutil.NewProviderFunc {
	_ = "STUB: not implemented"
	return *new(nodeutil.NewProviderFunc)
}

func virtRestConfig(ctx context.Context, virtualConfigPath string, hostClient ctrlruntimeclient.Client, clusterName, clusterNamespace, token string) (*rest.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// virtual kubeconfig file is empty, trying to fetch the k3k cluster kubeconfig

func kubeconfigBytes(url string, serverCA, clientCert, clientKey []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func addControllers(ctx context.Context, hostMgr, virtualMgr manager.Manager, c *config, hostClient ctrlruntimeclient.Client, virtEventRecorder record.EventRecorder) error {
	_ = "STUB: not implemented"
	return nil
}
