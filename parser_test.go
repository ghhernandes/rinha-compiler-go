package compiler_test

import (
	"github.com/ghhernandes/rinha-compiler-go"
	"os"
	"testing"
)

func BenchmarkParse(t *testing.B) {
	f, err := os.Open("files/fib.json")
	if err != nil {
		panic(err)
	}

	for i := 0; i < t.N; i++ {
		compiler.Parse(f)
	}
}
