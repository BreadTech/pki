package pkg

import (
	"crypto"
	"crypto/x509"
	"encoding/pem"
	"io"
	"os"
)

func PrintPrivateKeyPEM(alg string, key crypto.PrivateKey) error {
	return WritePrivateKeyPEM(os.Stdout, alg, key)
}

func WritePrivateKeyPEM(writer io.Writer, alg string, key crypto.PrivateKey) error {
	keyBytes, err := x509.MarshalPKCS8PrivateKey(key)
	if err != nil {
		return err
	}
	return pem.Encode(writer, &pem.Block{
		Type:    "PRIVATE KEY",
		Headers: map[string]string{"alg": alg},
		Bytes:   keyBytes,
	})
}
