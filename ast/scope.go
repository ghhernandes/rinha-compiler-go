package ast

type Scope map[string]Term

const SCOPE_DEFAULT_SIZE = 8

func (s Scope) Clone() Scope {
	clone := make(Scope, len(s)+SCOPE_DEFAULT_SIZE)
	for k, v := range s {
		clone[k] = v
	}
	return clone
}
