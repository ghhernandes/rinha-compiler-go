package compiler

import (
	"errors"
	"io"
)

type (
	Value any
	Scope map[string]Value
)

var (
	ErrTypeMismatch      = errors.New("type mismatch")
	ErrTypeNotComparable = errors.New("type not comparable")
	ErrInvalidBinaryOp   = errors.New("invalid binary operator")
)

func Eval(w io.Writer, f *File) error {
	t, err := eval(nil, f.Expression)
	if err != nil {
		return err
	}
	_, err = w.Write([]byte(stringify(t)))
	return err
}

func eval(scope Scope, t Term) (Term, error) {
	switch t.(type) {
	case Int, Str, Bool:
		return t, nil
	case Let:
		return let(scope, t.(Let))
	case Function:
		return function(scope, t.(Function))
	case If:
		return conditional(scope, t.(If))
	case Binary:
		return binary(t.(Binary))
	}
	return nil, nil
}

func binary(binary Binary) (Term, error) {
	switch binary.Op {
	case Eq:
		return eq(binary.Lhs, binary.Rhs)
	case Neq:
		return neq(binary.Lhs, binary.Rhs)
	case Lt:
		return lt(binary.Lhs, binary.Rhs)
	case Lte:
		return lte(binary.Lhs, binary.Rhs)
	case Gt:
		return gt(binary.Lhs, binary.Rhs)
	case Gte:
		return gte(binary.Lhs, binary.Rhs)
	case And:
		return and(binary.Lhs, binary.Rhs)
	case Or:
		return or(binary.Lhs, binary.Rhs)
	case Add:
		return add(binary.Lhs, binary.Rhs)
	case Sub:
		return sub(binary.Lhs, binary.Rhs)
	case Mul:
		return mul(binary.Lhs, binary.Rhs)
	case Div:
		return div(binary.Lhs, binary.Rhs)
	case Rem:
		return rem(binary.Lhs, binary.Rhs)
	default:
		return nil, ErrInvalidBinaryOp
	}
}

func eq(l, r Term) (Term, error) {
	switch l.(type) {
	case Int:
		return Bool{Kind: BOOL, Value: l.(Int).Value == r.(Int).Value}, nil
	case Str:
		return Bool{Kind: BOOL, Value: l.(Str).Value == r.(Str).Value}, nil
	case Bool:
		return Bool{Kind: BOOL, Value: l.(Bool).Value == r.(Bool).Value}, nil
	default:
		return nil, ErrTypeNotComparable
	}
}

func neq(l, r Term) (Term, error) {
	switch l.(type) {
	case Int:
		return Bool{Kind: BOOL, Value: l.(Int).Value != r.(Int).Value}, nil
	case Str:
		return Bool{Kind: BOOL, Value: l.(Str).Value != r.(Str).Value}, nil
	case Bool:
		return Bool{Kind: BOOL, Value: l.(Bool).Value != r.(Bool).Value}, nil
	default:
		return nil, ErrTypeNotComparable
	}
}

func lt(l, r Term) (Term, error) {
	switch l.(type) {
	case Int:
		return Bool{Kind: BOOL, Value: l.(Int).Value < r.(Int).Value}, nil
	case Str:
		return Bool{Kind: BOOL, Value: l.(Str).Value < r.(Str).Value}, nil
	default:
		return nil, ErrTypeNotComparable
	}
}

func lte(l, r Term) (Term, error) {
	switch l.(type) {
	case Int:
		return Bool{Kind: BOOL, Value: l.(Int).Value <= r.(Int).Value}, nil
	case Str:
		return Bool{Kind: BOOL, Value: l.(Str).Value <= r.(Str).Value}, nil
	default:
		return nil, ErrTypeNotComparable
	}
}

func gt(l, r Term) (Term, error) {
	switch l.(type) {
	case Int:
		return Bool{Kind: BOOL, Value: l.(Int).Value > r.(Int).Value}, nil
	case Str:
		return Bool{Kind: BOOL, Value: l.(Str).Value > r.(Str).Value}, nil
	default:
		return nil, ErrTypeNotComparable
	}
}

func gte(l, r Term) (Term, error) {
	switch l.(type) {
	case Int:
		return Bool{Kind: BOOL, Value: l.(Int).Value >= r.(Int).Value}, nil
	case Str:
		return Bool{Kind: BOOL, Value: l.(Str).Value >= r.(Str).Value}, nil
	default:
		return nil, ErrTypeNotComparable
	}
}

func and(l, r Term) (Term, error) {
	return Bool{Kind: BOOL, Value: l.(Bool).Value && r.(Bool).Value}, nil
}

func or(l, r Term) (Term, error) {
	return Bool{Kind: BOOL, Value: l.(Bool).Value || r.(Bool).Value}, nil
}

func add(l, r Term) (Term, error) {
	return Int{Kind: INT, Value: l.(Int).Value + r.(Int).Value}, nil
}

func sub(l, r Term) (Term, error) {
	return Int{Kind: INT, Value: l.(Int).Value - r.(Int).Value}, nil
}

func mul(l, r Term) (Term, error) {
	return Int{Kind: INT, Value: l.(Int).Value * r.(Int).Value}, nil
}

func div(l, r Term) (Term, error) {
	return Int{Kind: INT, Value: l.(Int).Value / r.(Int).Value}, nil
}

func rem(l, r Term) (Term, error) {
	return Int{Kind: INT, Value: l.(Int).Value % r.(Int).Value}, nil
}

func let(scope Scope, l Let) (Term, error) {
	value, err := eval(scope, l.Value)
	if err != nil {
		return nil, err
	}
	switch value.(type) {
	case Int, Str, Bool:
		scope[l.Name.Text] = value
	case Function:
		scope[l.Name.Text] = Function{Kind: FUNCTION, Parameters: value.(Function).Parameters, Value: value}
	}

	return eval(scope, l.Next)
}

func function(scope Scope, f Function) (Term, error) {
	return eval(scope, f.Value)
}

func conditional(scope Scope, cond If) (Term, error) {
	condition, err := eval(scope, cond.Condition)
	if err != nil {
		return nil, err
	}
	if condition.(Bool).Value {
		return eval(scope, cond.Then)
	}
	return eval(scope, cond.Otherwise)
}
