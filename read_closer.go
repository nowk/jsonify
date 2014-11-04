package jsonify

import (
	"io"
	"io/ioutil"
)

// NewReadCloser returns a marshalled struct as a ReadCloser
// Reader is obtained from NewReader
func NewReadCloser(d interface{}) (io.ReadCloser, error) {
	r, err := NewReader(d)
	if err != nil {
		return nil, err
	}

	return ioutil.NopCloser(r), nil
}
