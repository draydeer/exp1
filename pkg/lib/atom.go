package lib

import (
	"sync"
)

type Atom interface {
	Capture(key string) bool
	IsLocked(key string) bool
	Release(key string) Atom
}

type AtomInstance struct {
	sync.Mutex
	atoms map[string]bool
	isReleased chan string
}

func (atom *AtomInstance) Capture(key string) bool {
	defer atom.Unlock()

	atom.Lock()

	if _, isPresent := atom.atoms[key]; isPresent {
		return false
	}

	atom.atoms[key] = true

	return true
}

func (atom *AtomInstance) IsLocked(key string) bool {
	defer atom.Unlock()

	atom.Lock()

	_, isPresent := atom.atoms[key]

	return isPresent
}

func (atom *AtomInstance) Release(key string) Atom {
	defer atom.Unlock()

	atom.Lock()

	delete(atom.atoms, key)

	return atom
}

func NewAtom() *AtomInstance {
	return &AtomInstance{atoms: make(map[string]bool)}
}
