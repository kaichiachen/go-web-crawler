package reader

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
)

type MultipleReader interface {
	Reader() io.ReadCloser
}

type myMultipleReader struct {
	data []byte
}

func NewMultipleReader(reader io.Reader) (MultipleReader, error) {
	var data []byte
	var err error

	if reader != nil {
		data, err = ioutil.ReadAll(reader)
		if err != nil {
			return nil, fmt.Errorf("multiple reader: couldn't create a new one: %s", err)
		}
	} else {
		data = []byte{}
	}

	return &myMultipleReader{
		data: data,
	}, nil
}

func (rr *myMultipleReader) Reader() io.ReadCloser {
	return ioutil.NopCloser(bytes.NewReader(rr.data))
}
