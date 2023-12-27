package utils

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

var (
	ErrKeyMustBePEMEncoded = errors.New("key must be a PEM encoded PKCS1 or PKCS8 key")
	ErrNotRSAPublicKey     = errors.New("key is not a valid RSA public key")
)

// ParseRSAPublicKeyFromPEM parses a certificate or a PEM encoded PKCS1 or PKIX public key
func ParseRSAPublicKeyFromPEM(key []byte) (*rsa.PublicKey, error) {
	var err error

	// Parse PEM block
	var block *pem.Block
	if block, _ = pem.Decode(key); block == nil {
		return nil, ErrKeyMustBePEMEncoded
	}

	// Parse the key
	var parsedKey interface{}
	if parsedKey, err = x509.ParsePKIXPublicKey(block.Bytes); err != nil {
		if cert, err := x509.ParseCertificate(block.Bytes); err == nil {
			parsedKey = cert.PublicKey
		} else {
			if parsedKey, err = x509.ParsePKCS1PublicKey(block.Bytes); err != nil {
				return nil, err
			}
		}
	}

	var pkey *rsa.PublicKey
	var ok bool
	if pkey, ok = parsedKey.(*rsa.PublicKey); !ok {
		return nil, ErrNotRSAPublicKey
	}

	return pkey, nil
}
