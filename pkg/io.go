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
	buf := bytes.NewBuffer(dat)
	if _, err := buf.ReadFrom(reader); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
