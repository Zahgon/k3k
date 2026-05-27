package v1beta1

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"

	k3k "github.com/rancher/k3k/pkg/apis/k3k.io"
)

var (
	SchemeGroupVersion = schema.GroupVersion{Group: k3k.GroupName, Version: "v1beta1"}
	SchemBuilder       = runtime.NewSchemeBuilder(addKnownTypes)
	AddToScheme        = SchemBuilder.AddToScheme
)

func Resource(resource string) schema.GroupResource {
	_ = "STUB: not implemented"
	return *new(schema.GroupResource)
}

func addKnownTypes(s *runtime.Scheme) error { _ = "STUB: not implemented"; return nil }
