package vmd

import (
	"bytes"
	"io"
	"os"
)

type client struct {
	*reader
	*VMD
}

func Parse(path string) (*VMD, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	c := client{
		&reader{bytes.NewReader(data)},
		&VMD{},
	}

	return c.parse()
}
