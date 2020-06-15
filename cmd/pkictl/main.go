package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/BreadTech/pki/cmd/pkictl/generate"
	"github.com/BreadTech/pki/cmd/pkictl/hash"
	"github.com/BreadTech/pki/cmd/pkictl/sign"
)

func main() {
	cmd := &cobra.Command{
		Use:   "pkictl",
		Short: "For your public key infrastructure needs",
		Long:  "A command-line tool for administering public key infrastructure",
	}
	cmd.AddCommand(generate.Cmd)
	cmd.AddCommand(hash.Cmd)
	cmd.AddCommand(sign.Cmd)
	if err := cmd.Execute(); err != nil {
		logrus.WithError(err).Fatal("labelctl: error occurred")
	}
}
