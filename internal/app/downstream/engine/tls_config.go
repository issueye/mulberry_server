package engine

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"mulberry/internal/app/downstream/model"
	"os"
	"sync"
)

var (
	ErrCertNotFound = errors.New("certificate not found")
	certCache       = sync.Map{}
)

type TLSConfig struct {
	CertName string
	Cert     *tls.Certificate
	RootCAs  *x509.CertPool
}

func LoadCert(certInfo *model.CertInfo) (*tls.Certificate, error) {
	if certInfo.Public == "" || certInfo.Private == "" {
		return nil, ErrCertNotFound
	}

	// Check cache first
	if cert, ok := certCache.Load(certInfo.Name); ok {
		return cert.(*tls.Certificate), nil
	}

	// Load certificate
	cert, err := tls.LoadX509KeyPair(certInfo.Public, certInfo.Private)
	if err != nil {
		return nil, err
	}

	// Cache the certificate
	certCache.Store(certInfo.Name, &cert)
	return &cert, nil
}

func NewTLSConfig(certInfo *model.CertInfo) (*TLSConfig, error) {
	if certInfo == nil {
		return nil, nil
	}

	cert, err := LoadCert(certInfo)
	if err != nil {
		return nil, err
	}

	// Create root CA pool
	rootCAs := x509.NewCertPool()
	if caCert, err := os.ReadFile(certInfo.Public); err == nil {
		rootCAs.AppendCertsFromPEM(caCert)
	}

	return &TLSConfig{
		CertName: certInfo.Name,
		Cert:     cert,
		RootCAs:  rootCAs,
	}, nil
}
