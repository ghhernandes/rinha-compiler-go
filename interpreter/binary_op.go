package interpreter

import (
	"github.com/ghhernandes/rinha-compiler-go/ast"
)

func (i interpreter) eq(l, r ast.Term) ast.Term {
	switch n := l.(type) {
	case ast.Int:
		return ast.Bool{Kind: ast.BOOL, Value: n.Value == r.(ast.Int).Value}
	case ast.Str:
		return ast.Bool{Kind: ast.BOOL, Value: n.Value == r.(ast.Str).Value}
	case ast.Bool:
		return ast.Bool{Kind: ast.BOOL, Value: n.Value == r.(ast.Bool).Value}
	default:
		return nil
	}
}

func (i interpreter) neq(l, r ast.Term) ast.Term {
	switch l.(type) {
	case ast.Int:
		return ast.Bool{Kind: ast.BOOL, Value: l.(ast.Int).Value != r.(ast.Int).Value}
	case ast.Str:
		return ast.Bool{Kind: ast.BOOL, Value: l.(ast.Str).Value != r.(ast.Str).Value}
	case ast.Bool:
		return ast.Bool{Kind: ast.BOOL, Value: l.(ast.Bool).Value != r.(ast.Bool).Value}
	default:
		return nil
	}
}

func (i interpreter) lt(l, r ast.Term) ast.Term {
	switch l.(type) {
	case ast.Int:
		return ast.Bool{Kind: ast.BOOL, Value: l.(ast.Int).Value < r.(ast.Int).Value}
	case ast.Str:
		return ast.Bool{Kind: ast.BOOL, Value: l.(ast.Str).Value < r.(ast.Str).Value}
	default:
		return nil
	}
}

func (i interpreter) lte(l, r ast.Term) ast.Term {
	switch l.(type) {
	case ast.Int:
		return ast.Bool{Kind: ast.BOOL, Value: l.(ast.Int).Value <= r.(ast.Int).Value}
	case ast.Str:
		return ast.Bool{Kind: ast.BOOL, Value: l.(ast.Str).Value <= r.(ast.Str).Value}
	default:
		return nil
	}
}

func (i interpreter) gt(l, r ast.Term) ast.Term {
	switch l.(type) {
	case ast.Int:
		return ast.Bool{Kind: ast.BOOL, Value: l.(ast.Int).Value > r.(ast.Int).Value}
	case ast.Str:
		return ast.Bool{Kind: ast.BOOL, Value: l.(ast.Str).Value > r.(ast.Str).Value}
	default:
		return nil
	}
}

func (i interpreter) gte(l, r ast.Term) ast.Term {
	switch l.(type) {
	case ast.Int:
		return ast.Bool{Kind: ast.BOOL, Value: l.(ast.Int).Value >= r.(ast.Int).Value}
	case ast.Str:
		return ast.Bool{Kind: ast.BOOL, Value: l.(ast.Str).Value >= r.(ast.Str).Value}
	default:
		return nil
	}
}

func (i interpreter) and(l, r ast.Term) ast.Term {
	return ast.Bool{Kind: ast.BOOL, Value: l.(ast.Bool).Value && r.(ast.Bool).Value}
}

func (i interpreter) or(l, r ast.Term) ast.Term {
	return ast.Bool{Kind: ast.BOOL, Value: l.(ast.Bool).Value || r.(ast.Bool).Value}
}

func (i interpreter) add(l, r ast.Term) ast.Term {
	switch n := l.(type) {
	case ast.Int:
		return ast.Int{Kind: ast.INT, Value: n.Value + r.(ast.Int).Value}
	case ast.Str:
		return ast.Str{Kind: ast.STR, Value: n.Value + r.(ast.Str).Value}
	default:
		return nil
	}
}

func (i interpreter) sub(l, r ast.Term) ast.Term {
	return ast.Int{Kind: ast.INT, Value: l.(ast.Int).Value - r.(ast.Int).Value}
}

func (i interpreter) mul(l, r ast.Term) ast.Term {
	return ast.Int{Kind: ast.INT, Value: l.(ast.Int).Value * r.(ast.Int).Value}
}

func (i interpreter) div(l, r ast.Term) ast.Term {
	return ast.Int{Kind: ast.INT, Value: l.(ast.Int).Value / r.(ast.Int).Value}
}

func (i interpreter) rem(l, r ast.Term) ast.Term {
	return ast.Int{Kind: ast.INT, Value: l.(ast.Int).Value % r.(ast.Int).Value}
}
