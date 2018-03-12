package lib

import (
	"sync"
	"sync/atomic"
)

var transactionId uint64 = 0

type Atom interface {
	Capture(key string, transactionId uint64) bool
	IsCaptured(key string, transactionId uint64) bool
	Release(key string, transactionId uint64) Atom
}

type AtomTransactionInstance struct {
	sync.Mutex

	Id uint64
}

type AtomInstance struct {
	sync.Mutex

	Atoms map[string]*AtomTransactionInstance
}

func (atom *AtomInstance) Capture(key string, transactionId uint64) bool {
	atom.Lock()

	if transaction, isPresent := atom.Atoms[key]; isPresent {
		if transaction.Id == transactionId {
			atom.Unlock()

			return false
		}
	} else {
		atom.Atoms[key] = &AtomTransactionInstance{Id: transactionId}
	}

	atom.Unlock()

	atom.Atoms[key].Lock()

	atom.Atoms[key].Id = transactionId

	return true
}

func (atom *AtomInstance) IsCaptured(key string, transactionId uint64) bool {
	defer atom.Unlock()

	atom.Lock()

	if transaction, isPresent := atom.Atoms[key]; isPresent && transaction.Id == transactionId {
		return true
	}

	return false
}

func (atom *AtomInstance) Release(key string, transactionId uint64) Atom {
	defer atom.Unlock()

	atom.Lock()

	if transaction, isPresent := atom.Atoms[key]; isPresent && transaction.Id == transactionId {
		atom.Atoms[key].Unlock()
	}

	return atom
}

func NewAtom() *AtomInstance {
	return &AtomInstance{Atoms: make(map[string]*AtomTransactionInstance)}
}

func NewTransactionId() uint64 {
	return atomic.AddUint64(&transactionId, 1)
}
