package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/BreadTech/pki/cmd/pkictl/generate"
	"github.com/BreadTech/pki/cmd/pkictl/get"
	"github.com/BreadTech/pki/cmd/pkictl/hash"
	"github.com/BreadTech/pki/cmd/pkictl/read"
	"github.com/BreadTech/pki/cmd/pkictl/sign"
	"github.com/BreadTech/pki/cmd/pkictl/verify"

	bi "breadtech/interface"
)

type A struct {
	bi.Input
}

func main() {
	cmd := &cobra.Command{
		Use:   "pkictl",
		Short: "For your public key infrastructure needs",
		Long:  "A command-line tool for administering public key infrastructure",
	}
	cmd.AddCommand(generate.Cmd)
	cmd.AddCommand(get.Cmd)
	cmd.AddCommand(hash.Cmd)
	cmd.AddCommand(read.Cmd)
	cmd.AddCommand(sign.Cmd)
	cmd.AddCommand(verify.Cmd)
	if err := cmd.Execute(); err != nil {
		logrus.WithError(err).Fatal("labelctl: error occurred")
	}
}
