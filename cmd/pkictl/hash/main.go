package hash

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/BreadTech/pki/pkg"
)

var (
	Cmd = &cobra.Command{
		Aliases:   []string{"h", "dgst"},
		Args:      cobra.ExactArgs(1),
		ValidArgs: pkg.Hashes(),
		Use:       "hash",
		Short:     "Hash data",
		Run: func(cmd *cobra.Command, args []string) {
			dat, err := pkg.ReadFile(fileName)
			if err != nil {
				panic(err)
			}
			_, hashed := pkg.Hash(args[0], dat)
			switch outForm {
			case "hex":
				fmt.Println(hex.EncodeToString(hashed))
			case "b64":
				fmt.Println(base64.StdEncoding.EncodeToString(hashed))
			default:
				fmt.Println(string(hashed))
			}
		},
	}

	outForm  string
	fileName string
)

func init() {
	Cmd.Flags().StringVarP(&outForm, "output", "o", "raw", "Output form of data: [raw, hex, b64]")
	Cmd.Flags().StringVarP(&fileName, "file", "f", "", "Read data from file")
}
