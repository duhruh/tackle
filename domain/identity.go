package domain

type Identity interface {
	Identity() interface{}
}

type identity struct {
	id interface{}
}

func NewIdentity(id interface{}) Identity {
	return identity{id: id}
}

func (i identity) Identity() interface{} { return i.id }
