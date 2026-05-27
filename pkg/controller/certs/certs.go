package certs

import (
	"crypto/x509"
	"time"

	certutil "github.com/rancher/dynamiclistener/cert"
)

func CreateClientCertKey(commonName string, organization []string, altNames *certutil.AltNames, extKeyUsage []x509.ExtKeyUsage, expiresAt time.Duration, caCert, caKey string) ([]byte, []byte, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func generateKey() (data []byte, err error) { _ = "STUB: not implemented"; return nil, nil }

func AddSANs(sans []string) certutil.AltNames {
	_ = "STUB: not implemented"
	return *new(certutil.AltNames)
}
