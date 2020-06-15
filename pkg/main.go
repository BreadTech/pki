package pkg

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"syscall"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh/terminal"
)

var (
	PEMCiphers = map[string]x509.PEMCipher{
		"des":    x509.PEMCipherDES,
		"3des":   x509.PEMCipher3DES,
		"aes128": x509.PEMCipherAES128,
		"aes192": x509.PEMCipherAES192,
		"aes256": x509.PEMCipherAES256,
	}
)

// PrintPrivateKeyPEM encodes the given private key and writes to stdout.
func PrintPrivateKeyPEM(alg, cipher string, key crypto.PrivateKey) error {
	return WritePrivateKeyPEM(os.Stdout, alg, cipher, key)
}

// WritePrivateKeyPEM encodes the given private key and writes to given writer.
func WritePrivateKeyPEM(writer io.Writer, alg, cipher string, key crypto.PrivateKey) error {
	keyBytes, err := x509.MarshalPKCS8PrivateKey(key)
	if err != nil {
		return err
	}

	// If cipher is not none, then attempt to encrypt key.
	var block *pem.Block
	if cipher = strings.ToLower(cipher); cipher == "none" {
		block = &pem.Block{
			Type:  "PRIVATE KEY",
			Bytes: keyBytes,
		}
	} else {
		pwBytes, err := ReadSecureInput("Enter password: ")
		if err != nil {
			return err
		}

		pwBytes2, err := ReadSecureInput("Confirm password: ")
		if err != nil {
			return err
		}

		if !bytes.Equal(pwBytes, pwBytes2) {
			return fmt.Errorf("Passwords do not match")
		}

		cipher, ok := PEMCiphers[cipher]
		if !ok {
			return fmt.Errorf("Expected cipher: [des, 3des, aes128, aes192, aes256]")
		}

		if block, err = x509.EncryptPEMBlock(
			rand.Reader, fmt.Sprintf("%s PRIVATE KEY", strings.ToUpper(alg)),
			keyBytes, pwBytes, cipher); err != nil {
			return fmt.Errorf("Failed to encrypt key: %v", err)
		}
	}

	return pem.Encode(writer, block)
}

// PrintPublicKeyPEM encodes the given public key and writes to stdout.
func PrintPublicKeyPEM(alg string, key crypto.PublicKey) error {
	return WritePublicKeyPEM(os.Stdout, alg, key)
}

// WritePublicKeyPEM encodes the given public key and writes to given writer.
func WritePublicKeyPEM(writer io.Writer, alg string, key crypto.PublicKey) error {
	keyBytes, err := x509.MarshalPKIXPublicKey(key)
	if err != nil {
		return err
	}
	return pem.Encode(writer, &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: keyBytes,
	})
}

// ReadSecureInput reads input from stdin without echo.
func ReadSecureInput(prompt string) ([]byte, error) {
	fmt.Fprint(os.Stderr, prompt)
	pwBytes, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return nil, err
	}
	fmt.Fprintln(os.Stderr)
	return pwBytes, nil
}

// ReadPrivateKeyFile opens a file and parses the PEM bytes as PKCS8
func ReadPrivateKeyFile(fname string) (interface{}, error) {
	keyBytes, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}
	block, rest := pem.Decode(keyBytes)
	if len(rest) > 0 {
		logrus.WithField("rest", rest).Warn("pki/pkg.ReadKeyFile: extra data found in PEM")
	}
	// Prompt user for pw if encrypted.
	if block.Headers["DEK-Info"] != "" {
		pw, err := ReadSecureInput("Enter password: ")
		if err != nil {
			return nil, err
		}
		if block.Bytes, err = x509.DecryptPEMBlock(block, pw); err != nil {
			return nil, err
		}
	}
	return x509.ParsePKCS8PrivateKey(block.Bytes)
}
