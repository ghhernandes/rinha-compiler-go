package compiler_test

import (
	"github.com/ghhernandes/rinha-compiler-go"
	"github.com/ghhernandes/rinha-compiler-go/interpreter"
	"os"
	"testing"
)

func BenchmarkCompiler(t *testing.B) {

	f, err := os.Open("files/fib.json")
	if err != nil {
		panic(err)
	}

	ast, err := compiler.Parse(f)
	if err != nil {
		panic(err)
	}

	interpret := interpreter.New(nil, ast)

	for i := 0; i < t.N; i++ {
		if err := interpret.Execute(); err != nil {
			panic(err)
		}
	}
}
