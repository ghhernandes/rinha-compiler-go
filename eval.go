package compiler

import (
	"errors"
	"io"
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

func eval(env Env, t Term) (Term, error) {
	switch t.(type) {
	case Int, Str, Bool:
		return t, nil
	case Let:
		return t, nil
	case Binary:
		binary := t.(Binary)
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
	return nil, nil
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
