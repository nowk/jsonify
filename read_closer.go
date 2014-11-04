package jsonify

import (
	"io"
	"io/ioutil"
)

// NewReadCloser returns a marshalled struct as a ReadCloser
// Reader is from obtained from NewReader
func NewReadCloser(d interface{}) (io.ReadCloser, error) {
	r, err := NewReader(d)
	if err != nil {
		return nil, err
	}

	c := ioutil.NopCloser(r)
	return c, nil
}
