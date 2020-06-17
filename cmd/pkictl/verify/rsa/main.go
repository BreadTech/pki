package rsa

import (
	"crypto/rsa"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/BreadTech/pki/pkg"
)

const (
	keyType = "rsa"
)

var (
	hashType string
	datFile  string
	sigFile  string
	keyFile  string

	Cmd = &cobra.Command{
		Use:   keyType,
		Short: "Verify data using an RSA public key",
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

			// Read data to verify against.
			dat, err := pkg.ReadFile(datFile)
			if err != nil {
				panic(err)
			}

			// Verify
			hashEnum, hashed := pkg.Hash(hashType, dat)
			if err = rsa.VerifyPKCS1v15(
				key.(*rsa.PublicKey), hashEnum, hashed, sig,
			); err != nil {
				fmt.Println("Failed", err)
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
