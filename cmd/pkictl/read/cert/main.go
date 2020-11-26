package cert

import (
	//"crypto/x509"
	//"encoding/asn1"
	//"encoding/json"
	"encoding/pem"
	//"fmt"

	"github.com/spf13/cobra"
	//"gopkg.in/yaml.v2"

	"github.com/BreadTech/pki/pkg"
	//"github.com/BreadTech/pki/pkg/types"
	"github.com/BreadTech/pki/internal/asn/controller"
)

var (
	datFile string
	outForm string

	Cmd = &cobra.Command{
		Use:   "cert",
		Short: "Reads the data off a certificate",
		Run: func(cmd *cobra.Command, args []string) {
			// Read certificate file
			dat, err := pkg.ReadFile(datFile)
			if err != nil {
				panic(err)
			}

			/*
				printJSON := func(val *pkg.Node) error {
					out, err := json.Marshal(val)
					if err != nil {
						return err
					}
						indent := ""
						for i := 0; i < lvl; i++ {
							indent += "-"
						}

					fmt.Println(string(out))
					return nil
				}
			*/

			block, _ := pem.Decode(dat)

			svc, err := controller.New(block.Bytes)
			if err != nil {
				panic(err)
			}
			svc.Run()

			/*
				// Parse as cert
				cert, err := x509.ParseCertificate(block.Bytes)
				if err != nil {
					panic(err)
				}

				switch outForm {
				case "json":
					// Marshal as json
					out, err := json.Marshal(types.NewCert(cert))
					if err != nil {
						panic(err)
					}
					fmt.Println(string(out))
				default:
					// Marshal as yaml
					out, err := yaml.Marshal(types.NewCert(cert))
					if err != nil {
						panic(err)
					}
					fmt.Println(string(out))
				}
			*/

		},
	}
)

func init() {
	Cmd.Flags().StringVarP(&datFile, "file", "f", "", "Specifies the certificate file to read")
	Cmd.Flags().StringVarP(&outForm, "output", "o", "yaml", "Specifies the output format of how the certificate data is present")
	Cmd.MarkFlagRequired("file")
}
