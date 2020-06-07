package generate

import (
	"github.com/spf13/cobra"

	"github.com/BreadTech/pki/cmd/pkictl/generate/ecc"
	"github.com/BreadTech/pki/cmd/pkictl/generate/edc"
	"github.com/BreadTech/pki/cmd/pkictl/generate/rsa"
)

var (
	Cmd = &cobra.Command{
		Aliases: []string{"gen"},
		Use:     "generate",
		Short:   "Generate a key",
	}
)

func init() {
	Cmd.PersistentFlags().String("cipher", "aes128", "Specifies cipher to use for encrypting the private key")

	Cmd.AddCommand(ecc.Cmd)
	Cmd.AddCommand(edc.Cmd)
	Cmd.AddCommand(rsa.Cmd)
}
