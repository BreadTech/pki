package ec

import (
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/asn1"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"

	"github.com/spf13/cobra"

	"github.com/BreadTech/pki/pkg"
)

const (
	keyType = "ec"
)

var (
	hashType string
	datFile  string
	keyFile  string
	outForm  string

	Cmd = &cobra.Command{
		Use:   keyType,
		Short: "Signs data using an EC private key",
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
			_, hashed := pkg.Hash(hashType, dat)
			r, s, err := ecdsa.Sign(rand.Reader, key.(*ecdsa.PrivateKey), hashed)
			if err != nil {
				panic(err)
			}

			// Marshal the big ints as asn1
			if dat, err = asn1.Marshal([]*big.Int{r, s}); err != nil {
				panic(err)
			}

			switch outForm {
			case "raw":
				fmt.Print(string(dat))
			case "pem":
				pem.Encode(os.Stdout, &pem.Block{
					Type:  "EC SIGNATURE",
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
