package client

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// NewScheme creates a new Kubernetes runtime scheme with core APIs and k3k CRDs.
// This is suitable for most k3k test scenarios including integration and E2E tests.
func NewScheme() *runtime.Scheme { _ = "STUB: not implemented"; return nil }

// Add core Kubernetes scheme (includes most common types)

// Add k3k CRDs
