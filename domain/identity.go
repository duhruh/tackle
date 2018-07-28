package domain

import (
	"fmt"
)

// Identity -
type Identity interface {
	Identity() interface{}
}

type identity struct {
	id interface{}
}

// NewIdentity -
func NewIdentity(id interface{}) Identity {
	return identity{id: id}
}

func (i identity) Identity() interface{} { return i.id }
func (i identity) String() string {

	return fmt.Sprintf("%v", i.id)
}
