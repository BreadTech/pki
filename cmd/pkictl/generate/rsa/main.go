package rsa

import (
	"crypto/rand"
	"crypto/rsa"

	"github.com/spf13/cobra"

	"github.com/BreadTech/pki/pkg"
)

const (
	keyType = "rsa"
)

var (
	Cmd = &cobra.Command{
		Use:   keyType,
		Short: "Generates an RSA key",
		Run: func(cmd *cobra.Command, args []string) {
			privKey, err := rsa.GenerateKey(rand.Reader, 2048)
			if err != nil {
				panic(err)
			}

			if err = pkg.PrintPrivateKeyPEM(keyType, privKey); err != nil {
				panic(err)
			}
		},
	}
)
