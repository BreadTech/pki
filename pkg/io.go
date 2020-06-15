package pkg

import (
	"bytes"
	"io"
	"os"
)

func ReadFile(fileName string) ([]byte, error) {
	var reader io.Reader
	if fileName != "" {
		f, err := os.Open(fileName)
		if err != nil {
			return nil, err
		}
		defer f.Close()

		reader = f
	} else {
		reader = os.Stdin
	}
	dat := []byte{}
	if _, err := bytes.NewBuffer(dat).ReadFrom(reader); err != nil {
		return nil, err
	}
	return dat, nil
}
