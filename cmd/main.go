package main

import (
	"fmt"
	"os"

	"github.com/ghhernandes/rinha-compiler-go"
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

	if err := compiler.Eval(os.Stdout, ast); err != nil {
		panic(err)
	}
}
