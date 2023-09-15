package compiler

import (
	"encoding/json"
	"github.com/ghhernandes/rinha-compiler-go/ast"
	"io"
)

func Parse(r io.Reader) (*ast.File, error) {
	var f ast.File
	if err := json.NewDecoder(r).Decode(&f); err != nil {
		return nil, err
	}
	return &f, nil
}
