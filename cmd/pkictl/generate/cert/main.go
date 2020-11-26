package ec

import (
	"crypto/ecdsa"
	"crypto/rand"
	"strings"

	"github.com/spf13/cobra"

	"github.com/BreadTech/pki/pkg"
)

const (
	keyType = "ec"
)

var (
	selfSign bool
	keyFile  string
	certFile string

	Cmd = &cobra.Command{
		Use:   keyType,
		Short: "Generates a certificate",
		Long:  `Generates a certificate either as a signing request or self-signed`,
		Run: func(cmd *cobra.Command, args []string) {
			// Open private key
			key, err := pkg.ReadPrivateKeyFile(keyFile)
			if err != nil {
				panic(err)
			}

			// Open cert data file
			certDat, err := pkg.ReadFile(certFile)
			if err != nil {
				panic(err)
			}

		},
	}
)

func init() {
	Cmd.Flags().BoolVarP(&selfSign, "self-sign", "s", false, "Specifies to self-sign the certificate")
	Cmd.Flags().StringVarP(&keyFile, "key", "k", "", "Specifies the private key that owns the cert")
	Cmd.Flags().StringVarP(&certFile, "file", "f", "", "Specifies the data to place on cert")
}
