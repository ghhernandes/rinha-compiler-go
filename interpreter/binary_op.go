package interpreter

import (
	"strconv"
	"strings"

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
	switch n := l.(type) {
	case ast.Int:
		return ast.Bool{Kind: ast.BOOL, Value: n.Value != r.(ast.Int).Value}
	case ast.Str:
		return ast.Bool{Kind: ast.BOOL, Value: n.Value != r.(ast.Str).Value}
	case ast.Bool:
		return ast.Bool{Kind: ast.BOOL, Value: n.Value != r.(ast.Bool).Value}
	default:
		return nil
	}
}

func (i interpreter) lt(l, r ast.Term) ast.Term {
	switch n := l.(type) {
	case ast.Int:
		return ast.Bool{Kind: ast.BOOL, Value: n.Value < r.(ast.Int).Value}
	case ast.Str:
		return ast.Bool{Kind: ast.BOOL, Value: n.Value < r.(ast.Str).Value}
	default:
		return nil
	}
}

func (i interpreter) lte(l, r ast.Term) ast.Term {
	switch n := l.(type) {
	case ast.Int:
		return ast.Bool{Kind: ast.BOOL, Value: n.Value <= r.(ast.Int).Value}
	case ast.Str:
		return ast.Bool{Kind: ast.BOOL, Value: n.Value <= r.(ast.Str).Value}
	default:
		return nil
	}
}

func (i interpreter) gt(l, r ast.Term) ast.Term {
	switch n := l.(type) {
	case ast.Int:
		return ast.Bool{Kind: ast.BOOL, Value: n.Value > r.(ast.Int).Value}
	case ast.Str:
		return ast.Bool{Kind: ast.BOOL, Value: n.Value > r.(ast.Str).Value}
	default:
		return nil
	}
}

func (i interpreter) gte(l, r ast.Term) ast.Term {
	switch n := l.(type) {
	case ast.Int:
		return ast.Bool{Kind: ast.BOOL, Value: n.Value >= r.(ast.Int).Value}
	case ast.Str:
		return ast.Bool{Kind: ast.BOOL, Value: n.Value >= r.(ast.Str).Value}
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
	switch left := l.(type) {
	case ast.Int:
		switch right := r.(type) {
		case ast.Int:
			return ast.Int{Kind: ast.INT, Value: left.Value + right.Value}
		case ast.Str:
			return ast.Str{Kind: ast.STR, Value: strings.Join([]string{strconv.Itoa(int(left.Value)), right.Value}, "")}
		}
	case ast.Str:
		switch right := r.(type) {
		case ast.Int:
			return ast.Str{Kind: ast.STR, Value: strings.Join([]string{left.Value, strconv.Itoa(int(right.Value))}, "")}
		case ast.Str:
			return ast.Str{Kind: ast.STR, Value: left.Value + right.Value}
		}
	}
	return nil
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
