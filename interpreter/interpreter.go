package interpreter

import (
	"bytes"
	"fmt"
	"io"
	"strconv"

	"github.com/ghhernandes/rinha-compiler-go/ast"
	"github.com/ghhernandes/rinha-compiler-go/runtime"
)

const MEMOIZE_DELIMITER = ","

type interpreter struct {
	w   io.Writer
	f   *ast.File
	mem map[string]ast.Term
}

func New(w io.Writer, f *ast.File) *interpreter {
	return &interpreter{w: w, f: f, mem: make(map[string]ast.Term)}
}

func (i interpreter) Execute() error {
	scope := make(ast.Scope, ast.SCOPE_DEFAULT_SIZE)
	ast.Walk(i, scope, i.f.Expression)
	return nil
}

func (i interpreter) printTuple(b *bytes.Buffer, scope ast.Scope, t ast.Tuple) {
	printValueFn := func(node ast.Term) {
		switch n := node.(type) {
		case ast.Int:
			b.WriteString(strconv.FormatInt(int64(n.Value), 10))
		case ast.Str:
			b.WriteString(n.Value)
		case ast.Bool:
			b.WriteString(strconv.FormatBool(n.Value))
		case ast.Function:
			b.WriteString("<#closure>")
		default:
			b.WriteString("nil")
		}
	}

	first := i.eval(scope, t.First)
	second := i.eval(scope, t.Second)

	b.WriteString("(")
	printValueFn(first)
	b.WriteString(", ")
	printValueFn(second)
	b.WriteString(")")
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
	if i.w == nil {
		return nil
	}
	var b bytes.Buffer
	node := i.eval(scope, p.Value)
	switch n := node.(type) {
	case ast.Int:
		b.WriteString(strconv.FormatInt(int64(n.Value), 10))
	case ast.Str:
		b.WriteString(n.Value)
	case ast.Bool:
		b.WriteString(strconv.FormatBool(n.Value))
	case ast.Tuple:
		i.printTuple(&b, scope, n)
	case ast.Function:
		b.WriteString("<#closure>")
	default:
		b.WriteString("nil")
	}
	b.WriteString("\n")
	i.w.Write(b.Bytes())
	return nil
}

func (i interpreter) Call(scope ast.Scope, c ast.Call) ast.Term {
	callee := i.eval(scope, c.Callee)
	switch fn := callee.(type) {
	case ast.Function:
		if len(fn.Parameters) != len(c.Arguments) {
			runtime.Error(c.Location, "wrong number of arguments")
		}

		var b bytes.Buffer

		newScope := scope.Clone()
		if v, ok := c.Callee.(ast.Var); ok {
			b.WriteString(v.Text)
		} else {
			b.WriteString(strconv.FormatInt(int64(fn.Location.Start), 10))
			b.WriteString(MEMOIZE_DELIMITER)
			b.WriteString(strconv.FormatInt(int64(fn.Location.End), 10))
		}
		b.WriteString(MEMOIZE_DELIMITER)

		for index := 0; index < len(fn.Parameters); index++ {
			value := i.eval(scope, c.Arguments[index])
			newScope[fn.Parameters[index].Text] = value
			switch v := value.(type) {
			case ast.Int:
				b.WriteString(strconv.FormatInt(int64(v.Value), 10))
			case ast.Str:
				b.WriteString(v.Value)
			case ast.Bool:
				b.WriteString(strconv.FormatBool(v.Value))
			}
			b.WriteString(MEMOIZE_DELIMITER)
		}

		if memoized, ok := i.mem[b.String()]; ok {
			return memoized
		}
		evaluated := i.eval(newScope, fn.Value)
		i.mem[b.String()] = evaluated
		return evaluated
	default:
		return nil
	}
}

func (i interpreter) Tuple(scope ast.Scope, t ast.Tuple) ast.Term {
	return t
}

func (i interpreter) First(scope ast.Scope, f ast.First) ast.Term {
	node := i.eval(scope, f.Value)
	if tuple, ok := node.(ast.Tuple); ok {
		return i.eval(scope, tuple.First)
	}
	runtime.Error(f.Location, "not a tuple")
	return nil
}

func (i interpreter) Second(scope ast.Scope, s ast.Second) ast.Term {
	node := i.eval(scope, s.Value)
	if tuple, ok := node.(ast.Tuple); ok {
		return i.eval(scope, tuple.Second)
	}
	runtime.Error(s.Location, "not a tuple")
	return nil
}
