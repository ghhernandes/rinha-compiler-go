package ast

type Scope map[string]Term

func (s Scope) Clone() Scope {
	clone := make(Scope)
	for k, v := range s {
		clone[k] = v
	}
	return clone
}
