package ecc

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"strings"

	"github.com/spf13/cobra"

	"github.com/BreadTech/pki/pkg"
)

const (
	keyType = "ecc"
)

var (
	curveStr string
	curves   = map[string]elliptic.Curve{
		"p224": elliptic.P224(),
		"p256": elliptic.P256(),
		"p384": elliptic.P384(),
		"p521": elliptic.P521(),
	}

	Cmd = &cobra.Command{
		Use:   keyType,
		Short: "Generates an ellptic-curve key",
		Run: func(cmd *cobra.Command, args []string) {
			// Print public key for given private key
			if keyFile := cmd.Flag("public").Value.String(); keyFile != "" {
				privKey, err := pkg.ReadPrivateKeyFile(keyFile)
				if err != nil {
					panic(err)
				}
				eccKey, ok := privKey.(*ecdsa.PrivateKey)
				if !ok {
					panic("Not an elliptic-curve private key")
				}
				if err = pkg.PrintPublicKeyPEM(keyType, eccKey.Public()); err != nil {
					panic(err)
				}

				return
			}

			// Generate private key
			curve, ok := curves[strings.ToLower(curveStr)]
			if !ok {
				panic("curve must be one of [p224, p256, p384, p521]")
			}

			privKey, err := ecdsa.GenerateKey(curve, rand.Reader)
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
	Cmd.Flags().StringVarP(&curveStr, "curve", "c", "p224", "Specifies the curve to use")
}
