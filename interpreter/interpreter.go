package interpreter

import (
	"errors"
	"fmt"
	"github.com/ghhernandes/rinha-compiler-go/ast"
	"io"
)

type (
	Value any
	Scope map[string]Value
)

var (
	ErrTypeNotComparable = errors.New("type not comparable")
	ErrInvalidBinaryOp   = errors.New("invalid binary operator")
)

type interpreter struct {
	w io.Writer
	f *ast.File
}

func New(w io.Writer, f *ast.File) *interpreter {
	return &interpreter{w: w, f: f}
}

func (i interpreter) Execute() error {
	_, err := eval(make(Scope), i.f.Expression)
	if err != nil {
		return err
	}
	return err
}

func eval(scope Scope, t ast.Term) (ast.Term, error) {
	switch t.(type) {
	case ast.Int, ast.Str, ast.Bool:
		return t, nil
	case ast.Let:
		return let(scope, t.(ast.Let))
	case ast.Function:
		return function(scope, t.(ast.Function))
	case ast.If:
		return conditional(scope, t.(ast.If))
	case ast.Binary:
		return binary(t.(ast.Binary))
	case ast.Var:
		return var_(scope, t.(ast.Var))
	case ast.Print:
		return print_(scope, t.(ast.Print))
	case ast.Call:
		return call(scope, t.(ast.Call))
	}
	return nil, nil
}

func binary(binary ast.Binary) (ast.Term, error) {
	left, err := eval(nil, binary.Lhs)
	if err != nil {
		return nil, err
	}
	right, err := eval(nil, binary.Rhs)
	if err != nil {
		return nil, err
	}
	switch binary.Op {
	case ast.Eq:
		return eq(left, right)
	case ast.Neq:
		return neq(left, right)
	case ast.Lt:
		return lt(left, right)
	case ast.Lte:
		return lte(left, right)
	case ast.Gt:
		return gt(left, right)
	case ast.Gte:
		return gte(left, right)
	case ast.And:
		return and(left, right)
	case ast.Or:
		return or(left, right)
	case ast.Add:
		return add(left, right)
	case ast.Sub:
		return sub(left, right)
	case ast.Mul:
		return mul(left, right)
	case ast.Div:
		return div(left, right)
	case ast.Rem:
		return rem(left, right)
	default:
		return nil, ErrInvalidBinaryOp
	}
}

func eq(l, r ast.Term) (ast.Term, error) {
	switch l.(type) {
	case ast.Int:
		return ast.Bool{Kind: ast.BOOL, Value: l.(ast.Int).Value == r.(ast.Int).Value}, nil
	case ast.Str:
		return ast.Bool{Kind: ast.BOOL, Value: l.(ast.Str).Value == r.(ast.Str).Value}, nil
	case ast.Bool:
		return ast.Bool{Kind: ast.BOOL, Value: l.(ast.Bool).Value == r.(ast.Bool).Value}, nil
	default:
		return nil, ErrTypeNotComparable
	}
}

func neq(l, r ast.Term) (ast.Term, error) {
	switch l.(type) {
	case ast.Int:
		return ast.Bool{Kind: ast.BOOL, Value: l.(ast.Int).Value != r.(ast.Int).Value}, nil
	case ast.Str:
		return ast.Bool{Kind: ast.BOOL, Value: l.(ast.Str).Value != r.(ast.Str).Value}, nil
	case ast.Bool:
		return ast.Bool{Kind: ast.BOOL, Value: l.(ast.Bool).Value != r.(ast.Bool).Value}, nil
	default:
		return nil, ErrTypeNotComparable
	}
}

func lt(l, r ast.Term) (ast.Term, error) {
	switch l.(type) {
	case ast.Int:
		return ast.Bool{Kind: ast.BOOL, Value: l.(ast.Int).Value < r.(ast.Int).Value}, nil
	case ast.Str:
		return ast.Bool{Kind: ast.BOOL, Value: l.(ast.Str).Value < r.(ast.Str).Value}, nil
	default:
		return nil, ErrTypeNotComparable
	}
}

func lte(l, r ast.Term) (ast.Term, error) {
	switch l.(type) {
	case ast.Int:
		return ast.Bool{Kind: ast.BOOL, Value: l.(ast.Int).Value <= r.(ast.Int).Value}, nil
	case ast.Str:
		return ast.Bool{Kind: ast.BOOL, Value: l.(ast.Str).Value <= r.(ast.Str).Value}, nil
	default:
		return nil, ErrTypeNotComparable
	}
}

func gt(l, r ast.Term) (ast.Term, error) {
	switch l.(type) {
	case ast.Int:
		return ast.Bool{Kind: ast.BOOL, Value: l.(ast.Int).Value > r.(ast.Int).Value}, nil
	case ast.Str:
		return ast.Bool{Kind: ast.BOOL, Value: l.(ast.Str).Value > r.(ast.Str).Value}, nil
	default:
		return nil, ErrTypeNotComparable
	}
}

func gte(l, r ast.Term) (ast.Term, error) {
	switch l.(type) {
	case ast.Int:
		return ast.Bool{Kind: ast.BOOL, Value: l.(ast.Int).Value >= r.(ast.Int).Value}, nil
	case ast.Str:
		return ast.Bool{Kind: ast.BOOL, Value: l.(ast.Str).Value >= r.(ast.Str).Value}, nil
	default:
		return nil, ErrTypeNotComparable
	}
}

func and(l, r ast.Term) (ast.Term, error) {
	return ast.Bool{Kind: ast.BOOL, Value: l.(ast.Bool).Value && r.(ast.Bool).Value}, nil
}

func or(l, r ast.Term) (ast.Term, error) {
	return ast.Bool{Kind: ast.BOOL, Value: l.(ast.Bool).Value || r.(ast.Bool).Value}, nil
}

func add(l, r ast.Term) (ast.Term, error) {
	return ast.Int{Kind: ast.INT, Value: l.(ast.Int).Value + r.(ast.Int).Value}, nil
}

func sub(l, r ast.Term) (ast.Term, error) {
	return ast.Int{Kind: ast.INT, Value: l.(ast.Int).Value - r.(ast.Int).Value}, nil
}

func mul(l, r ast.Term) (ast.Term, error) {
	return ast.Int{Kind: ast.INT, Value: l.(ast.Int).Value * r.(ast.Int).Value}, nil
}

func div(l, r ast.Term) (ast.Term, error) {
	return ast.Int{Kind: ast.INT, Value: l.(ast.Int).Value / r.(ast.Int).Value}, nil
}

func rem(l, r ast.Term) (ast.Term, error) {
	return ast.Int{Kind: ast.INT, Value: l.(ast.Int).Value % r.(ast.Int).Value}, nil
}

func let(scope Scope, l ast.Let) (ast.Term, error) {
	value, err := eval(scope, l.Value)
	if err != nil {
		return nil, err
	}
	switch value.(type) {
	case ast.Int, ast.Str, ast.Bool:
		scope[l.Name.Text] = value
	case ast.Function:
		scope[l.Name.Text] = ast.Function{Kind: ast.FUNCTION, Parameters: value.(ast.Function).Parameters, Value: value}
	}

	return eval(scope, l.Next)
}

func function(scope Scope, f ast.Function) (ast.Term, error) {
	return f, nil
}

func conditional(scope Scope, cond ast.If) (ast.Term, error) {
	condition, err := eval(scope, cond.Condition)
	if err != nil {
		return nil, err
	}
	if condition.(ast.Bool).Value {
		return eval(scope, cond.Then)
	}
	return eval(scope, cond.Otherwise)
}

func var_(scope Scope, v ast.Var) (ast.Term, error) {
	return scope[v.Text], nil
}

func print_(scope Scope, p ast.Print) (ast.Term, error) {
	t, err := eval(scope, p.Value)
	if err != nil {
		return nil, err
	}

	switch t.(type) {
	case ast.Int:
		fmt.Println(t.(ast.Int).Value)
	case ast.Str:
		fmt.Println(t.(ast.Str).Value)
	case ast.Bool:
		fmt.Println(t.(ast.Bool).Value)
	}
	return nil, nil
}

func call(scope Scope, c ast.Call) (ast.Term, error) {
	return nil, nil
}
