package interpreter

import (
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/ghhernandes/rinha-compiler-go/ast"
	"github.com/ghhernandes/rinha-compiler-go/runtime"
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
	scope := make(map[string]ast.Term)
	ast.Walk(i, scope, i.f.Expression)
	return nil
}

func (i interpreter) eval(scope ast.Scope, expr ast.Term) ast.Term {
	return ast.Walk(i, scope, expr)
}

func (i interpreter) Bool(scope ast.Scope, b ast.Bool) ast.Term {
	return b
}

func (i interpreter) Int(scope ast.Scope, n ast.Int) ast.Term {
	return n
}

func (i interpreter) Str(scope ast.Scope, s ast.Str) ast.Term {
	return s
}

func (i interpreter) Binary(scope ast.Scope, binary ast.Binary) ast.Term {
	left := i.eval(scope, binary.Lhs)
	right := i.eval(scope, binary.Rhs)
	switch binary.Op {
	case ast.Eq:
		return i.eq(left, right)
	case ast.Neq:
		return i.neq(left, right)
	case ast.Lt:
		return i.lt(left, right)
	case ast.Lte:
		return i.lte(left, right)
	case ast.Gt:
		return i.gt(left, right)
	case ast.Gte:
		return i.gte(left, right)
	case ast.And:
		return i.and(left, right)
	case ast.Or:
		return i.or(left, right)
	case ast.Add:
		return i.add(left, right)
	case ast.Sub:
		return i.sub(left, right)
	case ast.Mul:
		return i.mul(left, right)
	case ast.Div:
		return i.div(left, right)
	case ast.Rem:
		return i.rem(left, right)
	default:
		return nil
	}
}

func (i interpreter) Let(scope ast.Scope, l ast.Let) ast.Term {
	scope[l.Name.Text] = i.eval(scope, l.Value)
	return i.eval(scope, l.Next)
}

func (i interpreter) Function(scope ast.Scope, f ast.Function) ast.Term {
	return f
}

func (i interpreter) If(scope ast.Scope, cond ast.If) ast.Term {
	condition := i.eval(scope, cond.Condition)
	if condition.(ast.Bool).Value {
		return i.eval(scope, cond.Then)
	}
	return i.eval(scope, cond.Otherwise)
}

func (i interpreter) Var(scope ast.Scope, v ast.Var) ast.Term {
	var (
		r  ast.Term
		ok bool
	)
	if r, ok = scope[v.Text]; !ok {
		runtime.Error(v.Location, fmt.Sprintf("undefined variable %s", v.Text))
	}
	return r
}

func (i interpreter) Print(scope ast.Scope, p ast.Print) ast.Term {
	node := i.eval(scope, p.Value)
	switch n := node.(type) {
	case ast.Int:
		i.w.Write([]byte(strconv.Itoa(int(n.Value))))
	case ast.Str:
		i.w.Write([]byte(n.Value))
	case ast.Bool:
		i.w.Write([]byte(strconv.FormatBool(n.Value)))
	}
	return nil
}

func (i interpreter) Call(scope ast.Scope, c ast.Call) ast.Term {
	callee := i.eval(scope, c.Callee)
	switch fn := callee.(type) {
	case ast.Function:
		if len(fn.Parameters) != len(c.Arguments) {
			runtime.Error(c.Location, "wrong number of arguments")
		}

		newScope := scope.Clone()
		for index := 0; index < len(fn.Parameters); index++ {
			newScope[fn.Parameters[index].Text] = i.eval(scope, c.Arguments[index])
		}

		return i.eval(newScope, fn.Value)
	}
	return nil
}
