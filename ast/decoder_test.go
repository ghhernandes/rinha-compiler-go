package ast_test

func BenchmarkDecoder(t *testing.B) {
	f, err := os.Open("../files/fib.json")
	if err != nil {
		panic(err)
	}

	ast, err := compiler.Parse(f)
	if err != nil {
		panic(err)
	}
}
