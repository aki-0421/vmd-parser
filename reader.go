package vmd

import (
	"bytes"
	"encoding/binary"
	"io"
	"math"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/transform"
)

// reader provides convenience functions for reading data from vmd files
type reader struct {
	*bytes.Reader
}

// ReadN returns the specified length of bytes
func (r reader) ReadN(n int) ([]byte, error) {
	b := make([]byte, n)

	_, err := r.Read(b)
	if err != nil {
		if err == io.EOF {
			err = nil
		}
		return b, err
	}

	return b, err
}

// GetStr returns the specified length of row charsets
func (r reader) GetChar(size int) (string, error) {
	res := make([]byte, 0, size)

	for i := 0; i < size; i++ {
		b, err := r.ReadN(1)
		if err != nil {
			return "", err
		}

		if b[0] == 0 {
			continue
		}

		res = append(res, b[0])
	}

	return string(res), nil
}

// GetUnicodeChar returns unicode charsets
func (r reader) GetUnicodeChar(size int) (string, error) {
	b, err := r.ReadN(size)
	if err != nil {
		return "", err
	}

	e, _ := charset.Lookup("shift_jis")
	u, _, err := transform.Bytes(
		e.NewDecoder(),
		b,
	)
	if err != nil {
		return "", err
	}

	return string(u), nil
}

// GetInt8 returns int8
func (r reader) GetInt8() (int8, error) {
	b, err := r.ReadN(1)
	if err != nil {
		return 0, err
	}

	return int8(b[0]), nil
}

// GetInt8Array returns int8 array
func (r reader) GetInt8Array(n int) ([]int8, error) {
	is := make([]int8, n)
	for i := 0; i < n; i++ {
		f, err := r.GetInt8()
		if err != nil {
			return is, err
		}

		is = append(is, f)
	}

	return is, nil
}

// GetInt returns int32
func (r reader) GetInt() (int, error) {
	b, err := r.ReadN(4)
	if err != nil {
		return 0, err
	}

	res := binary.LittleEndian.Uint32(b)

	return int(res), nil
}

// GetFloat32 returns float32
func (r reader) GetFloat32() (float32, error) {
	b, err := r.ReadN(4)
	if err != nil {
		return 0, err
	}

	bits := binary.LittleEndian.Uint32(b)
	res := math.Float32frombits(bits)

	return res, nil
}

// GetFloat32Array returns float32 array
func (r reader) GetFloat32Array(n int) ([]float32, error) {
	fs := make([]float32, n)
	for i := 0; i < n; i++ {
		f, err := r.GetFloat32()
		if err != nil {
			return fs, err
		}

		fs = append(fs, f)
	}

	return fs, nil
}

// newReader returns wrapped Reader
func newReader(r io.Reader) (*reader, error) {
	body, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	nr := bytes.NewReader(body)

	return &reader{nr}, nil
}
