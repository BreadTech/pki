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
			privKey, err := rsa.GenerateKey(rand.Reader, bitSize)
			if err != nil {
				panic(err)
			}

			if err = pkg.PrintPrivateKeyPEM(keyType, privKey); err != nil {
				panic(err)
			}
		},
	}
)

func init() {
	Cmd.Flags().IntVarP(&bitSize, "bit-size", "b", 2048, "Specifies the length of the key in bits")
}
