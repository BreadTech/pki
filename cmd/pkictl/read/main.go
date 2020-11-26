package read

import (
	"github.com/spf13/cobra"

	"github.com/BreadTech/pki/cmd/pkictl/read/cert"
)

var (
	Cmd = &cobra.Command{
		Aliases: []string{"r"},
		Use:     "read",
		Short:   "Decodes asn1 data",
	}
)

func init() {
	Cmd.AddCommand(cert.Cmd)
}
