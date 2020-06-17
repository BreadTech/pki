package ec

import (
	"crypto/ecdsa"
	"encoding/asn1"
	"encoding/pem"
	"fmt"
	"math/big"

	"github.com/spf13/cobra"

	"github.com/BreadTech/pki/pkg"
)

const (
	keyType = "ec"
)

var (
	hashType string
	datFile  string
	sigFile  string
	keyFile  string

	Cmd = &cobra.Command{
		Use:   keyType,
		Short: "Verify data using an EC public key",
		Run: func(cmd *cobra.Command, args []string) {
			// Open public key
			key, err := pkg.ReadPublicKeyFile(keyFile)
			if err != nil {
				panic(err)
			}

			// Read signature to verify.
			sig, err := pkg.ReadFile(sigFile)
			if err != nil {
				panic(err)
			}

			// Decode sig as pem if valid
			if pemBlock, _ := pem.Decode(sig); pemBlock != nil {
				sig = pemBlock.Bytes
			}

			// Decode sig as r, s
			ints := make([]*big.Int, 2)
			if _, err := asn1.Unmarshal(sig, &ints); err != nil {
				panic(err)
			}

			// Read data to verify against.
			dat, err := pkg.ReadFile(datFile)
			if err != nil {
				panic(err)
			}

			// Verify
			_, hashed := pkg.Hash(hashType, dat)
			if ok := ecdsa.Verify(key.(*ecdsa.PublicKey), hashed, ints[0], ints[1]); !ok {
				fmt.Println("Failed")
				return
			}
			fmt.Println("OK")
		},
	}
)

func init() {
	Cmd.Flags().StringVar(&hashType, "hash", "sha256", "Specifies the hash algorithm to use on the data")
	Cmd.Flags().StringVarP(&datFile, "file", "f", "", "Specifies the data file to verify against")
	Cmd.Flags().StringVarP(&keyFile, "key", "k", "", "Specifies the public key for verification")
	Cmd.Flags().StringVarP(&sigFile, "sig", "s", "", "Specifies the signature file to verify")
	Cmd.MarkFlagRequired("key")
}
