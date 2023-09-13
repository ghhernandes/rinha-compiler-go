package compiler

import (
	"encoding/json"
	"fmt"
)

func (f *File) UnmarshalJSON(data []byte) error {
	type Alias File
	aux := &struct {
		*Alias
		Expression json.RawMessage `json:"expression"`
	}{
		Alias: (*Alias)(f),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	var err error
	f.Expression, err = unmarshalTerm(aux.Expression)
	return err
}

func (op BinaryOp) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "Add":
		op = Add
	case "Sub":
		op = Sub
	case "Mul":
		op = Mul
	case "Div":
		op = Div
	case "Rem":
		op = Rem
	case "Eq":
		op = Eq
	case "Neq":
		op = Neq
	case "Lt":
		op = Lt
	case "Gt":
		op = Gt
	case "Lte":
		op = Lte
	case "Gte":
		op = Gte
	case "And":
		op = And
	case "Or":
		op = Or
	default:
		return fmt.Errorf("invalid binary operator: %s", s)
	}
	return nil
}

func unmarshalTerm(data []byte) (Term, error) {
	var t struct {
		Kind string `json:"kind"`
	}
	if err := json.Unmarshal(data, &t); err != nil {
		return nil, err
	}
	switch t.Kind {
	case INT:
		return unmarshalInt(data)
	case STR:
		return unmarshalStr(data)
	case BOOL:
		return unmarshalBool(data)
	case LET:
		return unmarshalLet(data)
	case VAR:
		return unmarshalVar(data)
	case FUNCTION:
		return unmarshalFunction(data)
	case CALL:
		return unmarshalCall(data)
	case IF:
		return unmarshalIf(data)
	case BINARY:
		return unmarshalBinary(data)
	case TUPLE:
		return unmarshalTuple(data)
	case PRINT:
		return unmarshalPrint(data)
	default:
		return nil, fmt.Errorf("invalid term kind: %s", t.Kind)
	}
}

func unmarshalInt(data []byte) (Int, error) {
	var i Int
	if err := json.Unmarshal(data, &i); err != nil {
		return Int{}, err
	}
	return i, nil
}

func unmarshalStr(data []byte) (Str, error) {
	var s Str
	if err := json.Unmarshal(data, &s); err != nil {
		return Str{}, err
	}
	return s, nil
}

func unmarshalBool(data []byte) (Bool, error) {
	var b Bool
	if err := json.Unmarshal(data, &b); err != nil {
		return Bool{}, err
	}
	return b, nil
}
func unmarshalLet(data []byte) (Let, error) {
	var (
		l   Let
		err error
	)
	if err := json.Unmarshal(data, &l); err != nil {
		return Let{}, err
	}

	bValue, err := json.Marshal(l.Value)
	if err != nil {
		return Let{}, err
	}
	bNext, err := json.Marshal(l.Next)
	if err != nil {
		return Let{}, err
	}

	l.Value, err = unmarshalTerm(bValue)
	if err != nil {
		return Let{}, err
	}
	l.Next, err = unmarshalTerm(bNext)
	return l, err
}

func unmarshalVar(data []byte) (Var, error) {
	var v Var
	if err := json.Unmarshal(data, &v); err != nil {
		return Var{}, err
	}
	return v, nil
}

func unmarshalFunction(data []byte) (Function, error) {
	var (
		f   Function
		err error
	)
	if err := json.Unmarshal(data, &f); err != nil {
		return Function{}, err
	}

	bValue, err := json.Marshal(f.Value)
	if err != nil {
		return Function{}, err
	}

	f.Value, err = unmarshalTerm(bValue)
	return f, err
}

func unmarshalCall(data []byte) (Call, error) {
	var (
		c   Call
		err error
	)

	if err := json.Unmarshal(data, &c); err != nil {
		return Call{}, err
	}

	bCallee, err := json.Marshal(c.Callee)
	if err != nil {
		return Call{}, err
	}

	bArgs, err := json.Marshal(c.Args)
	if err != nil {
		return Call{}, err
	}

	c.Callee, err = unmarshalTerm(bCallee)
	if err != nil {
		return Call{}, err
	}

	c.Args, err = unmarshalTerms(bArgs)
	return c, err
}

func unmarshalIf(data []byte) (If, error) {
	var (
		i   If
		err error
	)

	if err := json.Unmarshal(data, &i); err != nil {
		return If{}, err
	}

	bCondition, err := json.Marshal(i.Condition)
	if err != nil {
		return If{}, err
	}

	bThen, err := json.Marshal(i.Then)
	if err != nil {
		return If{}, err
	}

	bOtherwise, err := json.Marshal(i.Otherwise)
	if err != nil {
		return If{}, err
	}

	i.Condition, err = unmarshalTerm(bCondition)
	if err != nil {
		return If{}, err
	}

	i.Then, err = unmarshalTerm(bThen)
	if err != nil {
		return If{}, err
	}

	i.Otherwise, err = unmarshalTerm(bOtherwise)
	return i, err
}

func unmarshalBinary(data []byte) (Binary, error) {
	var (
		b   Binary
		err error
	)

	if err := json.Unmarshal(data, &b); err != nil {
		return Binary{}, err
	}

	bLhs, err := json.Marshal(b.Lhs)
	if err != nil {
		return Binary{}, err
	}

	bRhs, err := json.Marshal(b.Rhs)
	if err != nil {
		return Binary{}, err
	}

	b.Lhs, err = unmarshalTerm(bLhs)
	if err != nil {
		return Binary{}, err
	}

	b.Rhs, err = unmarshalTerm(bRhs)
	return b, err
}

func unmarshalTuple(data []byte) (Tuple, error) {
	var (
		t   Tuple
		err error
	)

	if err := json.Unmarshal(data, &t); err != nil {
		return Tuple{}, err
	}

	bFirst, err := json.Marshal(t.First)
	if err != nil {
		return Tuple{}, err
	}

	bSecond, err := json.Marshal(t.Second)
	if err != nil {
		return Tuple{}, err
	}

	t.First, err = unmarshalTerm(bFirst)
	if err != nil {
		return Tuple{}, err
	}

	t.Second, err = unmarshalTerm(bSecond)
	return t, err
}

func unmarshalPrint(data []byte) (Print, error) {
	var (
		p   Print
		err error
	)

	if err := json.Unmarshal(data, &p); err != nil {
		return Print{}, err
	}

	bValue, err := json.Marshal(p.Value)
	if err != nil {
		return Print{}, err
	}

	p.Value, err = unmarshalTerm(bValue)
	return p, err
}

func unmarshalTerms(data []byte) ([]Term, error) {
	var terms []Term
	if err := json.Unmarshal(data, &terms); err != nil {
		return nil, err
	}
	for i, term := range terms {
		bTerm, err := json.Marshal(term)
		if err != nil {
			return nil, err
		}
		terms[i], err = unmarshalTerm(bTerm)
		if err != nil {
			return nil, err
		}
	}
	return terms, nil
}
