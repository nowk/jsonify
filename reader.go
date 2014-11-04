package jsonify

import (
	"bytes"
	"encoding/json"
	"io"
)

// NewReader returns a marshalled struct as a Reader
func NewReader(d interface{}) (io.Reader, error) {
	b, err := json.Marshal(d)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(b), nil
}
