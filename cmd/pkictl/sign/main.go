package sign

import (
	"github.com/spf13/cobra"

	"github.com/BreadTech/pki/cmd/pkictl/sign/rsa"
)

var (
	Cmd = &cobra.Command{
		Aliases: []string{"sig", "s"},
		Use:     "sign",
		Short:   "Sign data using a private key",
	}
)

func init() {
	Cmd.AddCommand(rsa.Cmd)
}
