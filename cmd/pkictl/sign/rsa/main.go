package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/pem"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/BreadTech/pki/pkg"
)

const (
	keyType = "rsa"
)

var (
	hashType string
	datFile  string
	keyFile  string
	outForm  string

	Cmd = &cobra.Command{
		Use:   keyType,
		Short: "Signs data using an RSA private key",
		Run: func(cmd *cobra.Command, args []string) {
			// Open private key
			key, err := pkg.ReadPrivateKeyFile(keyFile)
			if err != nil {
				panic(err)
			}

			// Read data to hash
			dat, err := pkg.ReadFile(datFile)
			if err != nil {
				panic(err)
			}

			// Sign
			hashEnum, hashed := pkg.Hash(hashType, dat)
			if dat, err = rsa.SignPKCS1v15(
				rand.Reader, key.(*rsa.PrivateKey),
				hashEnum, hashed); err != nil {
				panic(err)
			}

			switch outForm {
			case "raw":
				fmt.Print(string(dat))
			case "pem":
				pem.Encode(os.Stdout, &pem.Block{
					Type:  "RSA SIGNATURE",
					Bytes: dat,
				})
			}
		},
	}
)

func init() {
	Cmd.Flags().StringVar(&hashType, "hash", "sha256", "Specifies the hash algorithm to use on the data")
	Cmd.Flags().StringVarP(&datFile, "file", "f", "", "Specifies the file to sign")
	Cmd.Flags().StringVarP(&keyFile, "key", "k", "", "Specifies the private key for signing")
	Cmd.Flags().StringVarP(&outForm, "out", "o", "raw", "Specifies the output format of the signature")
	Cmd.MarkFlagRequired("key")
}
