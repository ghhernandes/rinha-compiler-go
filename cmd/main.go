package main

import (
	"fmt"
	"os"

	"github.com/ghhernandes/rinha-compiler-go"
	"github.com/ghhernandes/rinha-compiler-go/interpreter"
)

func main() {
	f, err := os.Open("files/fib.json")
	if err != nil {
		panic(err)
	}

	ast, err := compiler.Parse(f)
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", ast))

	i := interpreter.New(os.Stdout, ast)
	if err := i.Execute(); err != nil {
		panic(err)
	}
}
