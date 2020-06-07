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
	bitSize int

	Cmd = &cobra.Command{
		Use:   keyType,
		Short: "Generates an RSA key",
		Run: func(cmd *cobra.Command, args []string) {
			// Print public key for given private key
			if keyFile := cmd.Flag("public").Value.String(); keyFile != "" {
				privKey, err := pkg.ReadPrivateKeyFile(keyFile)
				if err != nil {
					panic(err)
				}
				rsaKey, ok := privKey.(*rsa.PrivateKey)
				if !ok {
					panic("Not an RSA private key")
				}
				if err = pkg.PrintPublicKeyPEM(keyType, rsaKey.Public()); err != nil {
					panic(err)
				}

				return
			}

			// Generate private key
			privKey, err := rsa.GenerateKey(rand.Reader, bitSize)
			if err != nil {
				panic(err)
			}

			if err = pkg.PrintPrivateKeyPEM(keyType, cmd.Flag("cipher").Value.String(), privKey); err != nil {
				panic(err)
			}
		},
	}
)

func init() {
	Cmd.Flags().IntVarP(&bitSize, "bit-size", "b", 2048, "Specifies the length of the key in bits")
}
