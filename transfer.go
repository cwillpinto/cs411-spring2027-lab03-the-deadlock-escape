package transfer

import (
	"runtime"
	"sync"
)

// Account IDs are unique. Balance is protected by mu.
type Account struct {
	ID      int
	mu      sync.Mutex
	balance int
}

func NewAccount(id, opening int) *Account {
	return &Account{ID: id, balance: opening}
}

func (a *Account) Balance() int {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.balance
}

// Transfer moves amount from one account to another. It returns false if the
// source lacks funds. TODO: replace caller-order acquisition with one global
// order based on the unique IDs. Unlock in reverse order; keep both balance
// changes inside the critical section.
func Transfer(from, to *Account, amount int) bool {
	if from == nil || to == nil || amount < 0 || from == to {
		return false
	}
	from.mu.Lock()
	defer from.mu.Unlock()
	runtime.Gosched() // widen the window to make the lock-order bug visible
	to.mu.Lock()
	defer to.mu.Unlock()
	if from.balance < amount {
		return false
	}
	from.balance -= amount
	to.balance += amount
	return true
}
