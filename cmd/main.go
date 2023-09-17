package main

import (
	"io"
	"os"

	"github.com/ghhernandes/rinha-compiler-go"
	"github.com/ghhernandes/rinha-compiler-go/interpreter"
)

func main() {
	var (
		f   io.Reader
		err error
	)

	if len(os.Args) > 1 {
		f, err = os.Open(os.Args[1])
		if err != nil {
			panic(err)
		}
	} else {
		f = os.Stdin
	}

	ast, err := compiler.Parse(f)
	if err != nil {
		panic(err)
	}

	if err := interpreter.New(os.Stdout, ast).Execute(); err != nil {
		panic(err)
	}
}
