package mounts

import (
	corev1 "k8s.io/api/core/v1"

	"github.com/rancher/k3k/pkg/apis/k3k.io/v1beta1"
)

func BuildSecretsMountsVolumes(secretMounts []v1beta1.SecretMount, role string) ([]corev1.Volume, []corev1.VolumeMount) {
	_ = "STUB: not implemented"
	return nil, nil
}

func buildSecretMountVolume(secretMount v1beta1.SecretMount) (corev1.Volume, corev1.VolumeMount) {
	_ = "STUB: not implemented"
	return *new(corev1.Volume), *new(corev1.VolumeMount)
}
