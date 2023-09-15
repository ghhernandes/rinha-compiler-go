package ast

import (
	"fmt"
)

type Error struct {
	error
	loc Location
}

func NewError(loc Location, err error) Error {
	return Error{
		error: err,
		loc:   loc,
	}
}

func (e Error) Error() string {
	return fmt.Sprintf("%s:%d:%d: %s", e.loc.Filename, e.loc.Start, e.loc.End, e.error.Error())
}
