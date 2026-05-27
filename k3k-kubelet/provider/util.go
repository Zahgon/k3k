package provider

import (
	"github.com/virtual-kubelet/virtual-kubelet/node/api"
	"k8s.io/client-go/tools/remotecommand"
)

// translatorSizeQueue feeds the size events from the WebSocket
// resizeChan into the SPDY client input. Implements TerminalSizeQueue
// interface.
type translatorSizeQueue struct {
	resizeChan <-chan api.TermSize
}

func (t *translatorSizeQueue) Next() *remotecommand.TerminalSize {
	_ = "STUB: not implemented"
	return nil
}
