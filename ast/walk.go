package ast

type Visitor interface {
	Int(Scope, Int) Term
	Str(Scope, Str) Term
	Bool(Scope, Bool) Term
	Let(Scope, Let) Term
	Function(Scope, Function) Term
	If(Scope, If) Term
	Binary(Scope, Binary) Term
	Var(Scope, Var) Term
	Print(Scope, Print) Term
	Call(Scope, Call) Term
}

func Walk(v Visitor, scope Scope, node Term) Term {
	switch n := node.(type) {
	case Int, Str, Bool:
		return n
	case Let:
		return v.Let(scope, n)
	case Function:
		return v.Function(scope, n)
	case If:
		return v.If(scope, n)
	case Binary:
		return v.Binary(scope, n)
	case Var:
		return v.Var(scope, n)
	case Print:
		return v.Print(scope, n)
	case Call:
		return v.Call(scope, n)
	}
	return nil
}
