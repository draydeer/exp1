package lib

type Atom interface {
	Capture(key string) bool
	IsLocked(key string) bool
	Release(key string) Atom
}

type AtomInstance struct {
	atoms map[string]bool
}

func (atom *AtomInstance) Capture(key string) bool {
	if atom.IsLocked(key) {
		return false
	}

	atom.atoms[key] = true

	return true
}

func (atom *AtomInstance) IsLocked(key string) bool {
	_, isPresent := atom.atoms[key]

	return isPresent
}

func (atom *AtomInstance) Release(key string) Atom {
	delete(atom.atoms, key)

	return atom
}

func NewAtom() *AtomInstance {
	return &AtomInstance{make(map[string]bool)}
}
