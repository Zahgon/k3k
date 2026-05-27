package k3k

import (
	"k8s.io/client-go/kubernetes"

	corev1 "k8s.io/api/core/v1"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// CreateNamespace creates a new namespace with a generated name and the "e2e: true" label.
// The namespace is created using the provided Kubernetes clientset.
func CreateNamespace(clientset kubernetes.Interface) *corev1.Namespace {
	_ = "STUB: not implemented"
	return nil
}

// DeleteNamespaces deletes the specified namespaces in parallel.
// If the KEEP_NAMESPACES environment variable is set, namespaces are preserved instead.
// This is useful for debugging test failures.
func DeleteNamespaces(clientset kubernetes.Interface, names ...string) {
	_ = "STUB: not implemented"
	return
}
