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
			_, privKey, err := ed25519.GenerateKey(rand.Reader)
			if err != nil {
				panic(err)
			}

			if err = pkg.PrintPrivateKeyPEM(keyType, privKey); err != nil {
				panic(err)
			}
		},
	}
)
