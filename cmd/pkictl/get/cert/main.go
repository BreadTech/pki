package cert

import (
	"crypto/tls"
	"encoding/pem"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	Cmd = &cobra.Command{
		Use:   "cert",
		Short: "Retrieves the certificate(s) of a TLS connection",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			conn, err := tls.Dial("tcp", args[0], &tls.Config{InsecureSkipVerify: true})
			if err != nil {
				panic(err)
			}

			state := conn.ConnectionState()
			hasChain := len(state.VerifiedChains) > 0
			for i, cert := range state.PeerCertificates {
				pem.Encode(os.Stdout, &pem.Block{
					Type:  "CERTIFICATE",
					Bytes: cert.Raw,
				})
				fmt.Println()
				if hasChain {
					for _, chainCert := range state.VerifiedChains[i] {
						pem.Encode(os.Stdout, &pem.Block{
							Type:  "CERTIFICATE",
							Bytes: chainCert.Raw,
						})
						fmt.Println()
					}
				}
			}
		},
	}
)
