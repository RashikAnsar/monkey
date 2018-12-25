package object

import (
	"fmt"
)

type ObjectType string

const (
	INTEGER_OBJ = "INTEGER"
)

// Object interface with Type and Inspect Methods
type Object interface {
	Type() ObjectType
	Inspect() string
}

// Integer type of monkey language
type Integer struct {
	Value int64
}

func (i *Integer) Type() ObjectType { return INTEGER_OBJ }
func (i *Integer) Inspect() string  { return fmt.Sprintf("%d", i.Value) }
