package interpreter_test

import (
	"github.com/ghhernandes/rinha-compiler-go"
	"github.com/ghhernandes/rinha-compiler-go/interpreter"
	"os"
	"testing"
)

func TestInterpreter(t *testing.T) {
	f, err := os.Open("../files/fib.json")
	if err != nil {
		panic(err)
	}

	ast, err := compiler.Parse(f)
	if err != nil {
		panic(err)
	}

	interpret := interpreter.New(nil, ast)
	if err := interpret.Execute(); err != nil {
		t.Fail()
	}
}

func BenchmarkInterpreter(t *testing.B) {
	f, err := os.Open("../files/fib.json")
	if err != nil {
		panic(err)
	}

	ast, err := compiler.Parse(f)
	if err != nil {
		panic(err)
	}

	interpret := interpreter.New(nil, ast)
	for i := 0; i < t.N; i++ {
		interpret.Execute()
	}
}
