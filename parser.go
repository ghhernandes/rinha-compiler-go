package compiler

import (
	"encoding/json"
	"io"
)

func Parse(r io.Reader) (*File, error) {
	var f File
	if err := json.NewDecoder(r).Decode(&f); err != nil {
		return nil, err
	}
	return &f, nil
}
