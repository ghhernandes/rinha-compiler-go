package runtime

import (
	"fmt"
	"github.com/ghhernandes/rinha-compiler-go/ast"
)

func Error(loc ast.Location, msg string) {
	panic(fmt.Errorf("%s:%d:%d: %s", loc.Filename, loc.Start, loc.End, msg))
}
