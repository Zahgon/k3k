package main

import (
	"crypto/tls"
	"net/http"

	"github.com/virtual-kubelet/virtual-kubelet/node/nodeutil"
)

func (k *kubelet) registerNode(agentIP, podIP string, cfg config) error {
	_ = "STUB: not implemented"
	return nil
}

func nodeOpt(mux *http.ServeMux, tlsConfig *tls.Config, port int) nodeutil.NodeOpt {
	_ = "STUB: not implemented"
	return *new(nodeutil.NodeOpt)
}

func loadTLSConfig(cfg config, nodeName, token, agentIP, podIP string) (*tls.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// create rootCA CertPool
