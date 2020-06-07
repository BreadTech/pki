package pkg

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io"
	"os"
	"strings"
	"syscall"

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
			Type:    "PRIVATE KEY",
			Headers: map[string]string{"alg": alg},
			Bytes:   keyBytes,
		}
	} else {
		pwBytes, err := ReadSecureInput("Enter Password: ")
		if err != nil {
			return err
		}

		pwBytes2, err := ReadSecureInput("Confirm Password: ")
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
			rand.Reader, "ENCRYPTED PRIVATE KEY",
			keyBytes, pwBytes, cipher); err != nil {
			return err
		}
	}

	return pem.Encode(writer, block)
}

// ReadSecureInput reads input from stdin without echo.
func ReadSecureInput(prompt string) ([]byte, error) {
	fmt.Print(prompt)
	pwBytes, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return nil, err
	}
	fmt.Println()
	return pwBytes, nil
}
