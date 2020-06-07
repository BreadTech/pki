package edc

import (
	"crypto/ed25519"
	"crypto/rand"

	"github.com/spf13/cobra"

	"github.com/BreadTech/pki/pkg"
)

const (
	keyType = "edc"
)

var (
	Cmd = &cobra.Command{
		Use:   "edc",
		Short: "Generates an ed25519 key",
		Run: func(cmd *cobra.Command, args []string) {
			// Print public key for given private key
			if keyFile := cmd.Flag("public").Value.String(); keyFile != "" {
				privKey, err := pkg.ReadPrivateKeyFile(keyFile)
				if err != nil {
					panic(err)
				}
				edcKey, ok := privKey.(ed25519.PrivateKey)
				if !ok {
					panic("Not an Edwards-curve private key")
				}
				if err = pkg.PrintPublicKeyPEM(keyType, edcKey.Public()); err != nil {
					panic(err)
				}

				return
			}

			// Generate private key
			_, privKey, err := ed25519.GenerateKey(rand.Reader)
			if err != nil {
				panic(err)
			}

			if err = pkg.PrintPrivateKeyPEM(keyType, cmd.Flag("cipher").Value.String(), privKey); err != nil {
				panic(err)
			}
		},
	}
)
