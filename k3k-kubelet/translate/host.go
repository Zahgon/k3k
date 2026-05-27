package translate

import (
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	corev1 "k8s.io/api/core/v1"

	"github.com/rancher/k3k/pkg/apis/k3k.io/v1beta1"
)

const (
	// ClusterNameLabel is the key for the label that contains the name of the virtual cluster
	// this resource was made in
	ClusterNameLabel = "k3k.io/clusterName"
	// ResourceNameAnnotation is the key for the annotation that contains the original name of this
	// resource in the virtual cluster
	ResourceNameAnnotation = "k3k.io/name"
	// ResourceNamespaceAnnotation is the key for the annotation that contains the original namespace of this
	// resource in the virtual cluster
	ResourceNamespaceAnnotation = "k3k.io/namespace"
	// MetadataNameField is the downwardapi field for object's name
	MetadataNameField = "metadata.name"
	// MetadataNamespaceField is the downward field for the object's namespace
	MetadataNamespaceField = "metadata.namespace"
)

type ToHostTranslator struct {
	// ClusterName is the name of the virtual cluster whose resources we are
	// translating to a host cluster
	ClusterName string
	// ClusterNamespace is the namespace of the virtual cluster whose resources
	// we are translating to a host cluster
	ClusterNamespace string
}

func NewHostTranslator(cluster *v1beta1.Cluster) *ToHostTranslator {
	_ = "STUB: not implemented"
	return nil
}

// Translate translates a virtual cluster object to a host cluster object. This should only be used for
// static resources such as configmaps/secrets, and not for things like pods (which can reference other
// objects). Note that this won't set host-cluster values (like resource version) so when updating you
// may need to fetch the existing value and do some combination before using this.
func (t *ToHostTranslator) TranslateTo(obj client.Object) {
	_ = "STUB: not implemented"
	// owning objects may be in the virtual cluster, but may not be in the host cluster
	return
}

// add some annotations to make it easier to track source object

// add a label to quickly identify objects owned by a given virtual cluster

// resource version/UID won't match what's in the host cluster.

// set the name and the namespace so that this goes in the proper host namespace
// and doesn't collide with other resources

func (t *ToHostTranslator) TranslateFrom(obj client.Object) {
	_ = "STUB: not implemented"
	// owning objects may be in the virtual cluster, but may not be in the host cluster
	return
}

// remove the annotations added to track original name

// TODO: It's possible that this was erased by a change on the host cluster
// In this case, we need to have some sort of fallback or error return

// remove the clusteName tracking label

// resource version/UID won't match what's in the virtual cluster.

// TranslateName returns the name of the resource in the host cluster. Will not update the object with this name.
func (t *ToHostTranslator) TranslateName(namespace string, name string) string {
	_ = "STUB: not implemented"

	// some resources are not namespaced (i.e. priorityclasses)
	/// for these resources we skip the namespace to avoid having a name like: prioritclass--cluster-123
	return ""
}

// we need to come up with a name which is:
// - somewhat connectable to the original resource
// - a valid k8s name
// - idempotently calculatable
// - unique for this combination of name/namespace/cluster

// use + as a separator since it can't be in an object name

// it's possible that the suffix will be in the name, so we use hex to make it valid for k8s

// NamespacedName returns the types.NamespacedName of the resource in the host cluster
func (t *ToHostTranslator) NamespacedName(obj client.Object) types.NamespacedName {
	_ = "STUB: not implemented"
	return *new(types.NamespacedName)
}

// TranslateObjectReferenceFrom translates a host-cluster ObjectReference back to
// virtual-cluster coordinates by reversing the name encoding applied by TranslateName.
// If the name cannot be reversed (e.g. it was truncated by SafeConcatName), the
// original host name and namespace are preserved.
func (t *ToHostTranslator) TranslateObjectReferenceFrom(ref corev1.ObjectReference) *corev1.ObjectReference {
	_ = "STUB: not implemented"
	return nil
}

// reverseTranslateName attempts to recover the original virtual name and namespace
// from a host-cluster translated name. TranslateName encodes the original values as
// a hex string suffix (hex("name+namespace+clusterName")), which this method decodes.
// Returns ok=false when the name was truncated and cannot be reversed.
func (t *ToHostTranslator) reverseTranslateName(translatedName string) (name, namespace string, ok bool) {
	_ = "STUB: not implemented"
	return "", "", false
}
