package controller

import (
	"time"

	"k8s.io/apimachinery/pkg/util/wait"

	"github.com/rancher/k3k/pkg/apis/k3k.io/v1beta1"
)

const (
	namePrefix      = "k3k"
	AdminCommonName = "system:admin"
)

// Backoff is the cluster creation duration backoff
var Backoff = wait.Backoff{
	Steps:    5,
	Duration: 5 * time.Second,
	Factor:   2,
	Jitter:   0.1,
}

// K3SImage returns the rancher/k3s image tagged with the found K3SVersion.
func K3SImage(cluster *v1beta1.Cluster, k3SImage string) string {
	_ = "STUB: not implemented"
	return ""
}

// K3SVersion returns the rancher/k3s specified version.
// If empty it will return the k3s version of the Kubernetes version of the host cluster, stored in the Status object.
// Returns the latest version as fallback.
func K3SVersion(cluster *v1beta1.Cluster) string { _ = "STUB: not implemented"; return "" }

// SafeConcatNameWithPrefix runs the SafeConcatName with extra prefix.
func SafeConcatNameWithPrefix(name ...string) string { _ = "STUB: not implemented"; return "" }

// SafeConcatName concatenates the given strings and ensures the returned name is under 64 characters
// by cutting the string off at 57 characters and setting the last 6 with an encoded version of the concatenated string.
// Empty strings in the array will be ignored.
func SafeConcatName(name ...string) string { _ = "STUB: not implemented"; return "" }

// since we cut the string in the middle, the last char may not be compatible with what is expected in k8s
// we are checking and if necessary removing the last char
