package ast

type Scope map[string]Term

const SCOPE_DEFAULT_SIZE = 32

func (s Scope) Clone() Scope {
	clone := make(Scope, len(s)+SCOPE_DEFAULT_SIZE)
	for k, v := range s {
		clone[k] = v
	}
	return clone
}

func (s Scope) Zip(other Scope) Scope {
    clone := make(Scope, len(s)+len(other))
    for k, v := range s {
        clone[k] = v
    }
    for k, v := range other {
        clone[k] = v
    }
    return clone
}
