package ecc

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"

	"github.com/spf13/cobra"

	"github.com/BreadTech/pki/pkg"
)

const (
	keyType = "ecc"
)

var (
	curves = map[string]elliptic.Curve{
		"p224": elliptic.P224(),
		"p256": elliptic.P256(),
		"p384": elliptic.P384(),
		"p521": elliptic.P521(),
	}

	Cmd = &cobra.Command{
		Use:   keyType,
		Short: "Generates an ellptic-curve key",
		Run: func(cmd *cobra.Command, args []string) {
			privKey, err := ecdsa.GenerateKey(curves["p224"], rand.Reader)
			if err != nil {
				panic(err)
			}

			if err = pkg.PrintPrivateKeyPEM(keyType, privKey); err != nil {
				panic(err)
			}

		},
	}
)
