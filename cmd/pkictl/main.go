package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/breadtech/pki/cmd/pkictl/generate"
)

func main() {
	cmd := &cobra.Command{
		Use:   "pkictl",
		Short: "For your public key infrastructure needs",
		Long:  "A command-line tool for administering public key infrastructure",
	}
	cmd.AddCommand(generate.Cmd)
	if err := cmd.Execute(); err != nil {
		logrus.WithError(err).Fatal("pkictl: error occurred")
	}
}
