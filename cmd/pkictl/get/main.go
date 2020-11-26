package get

import (
	"github.com/spf13/cobra"

	"github.com/BreadTech/pki/cmd/pkictl/get/cert"
)

var (
	Cmd = &cobra.Command{
		Aliases: []string{"g"},
		Use:     "get",
		Short:   "Retrieves data from a remote server",
	}
)

func init() {
	Cmd.AddCommand(cert.Cmd)
}
