package verify

import (
	"github.com/spf13/cobra"

	"github.com/BreadTech/pki/cmd/pkictl/verify/ec"
	"github.com/BreadTech/pki/cmd/pkictl/verify/rsa"
)

var (
	Cmd = &cobra.Command{
		Aliases: []string{"v"},
		Use:     "verify",
		Short:   "Verify a signature using a public key",
	}
)

func init() {
	Cmd.AddCommand(rsa.Cmd)
	Cmd.AddCommand(ec.Cmd)
}
