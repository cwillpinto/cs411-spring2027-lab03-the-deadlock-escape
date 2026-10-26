package transfer

import (
	"sync"
	"testing"
)

func TestTransfersSequential(t *testing.T) {
	a, b := NewAccount(1, 100), NewAccount(2, 100)
	if !Transfer(a, b, 25) {
		t.Fatal("transfer failed")
	}
	if a.Balance() != 75 || b.Balance() != 125 {
		t.Fatal("wrong balances")
	}
	if Transfer(a, b, 100) {
		t.Fatal("overdraft accepted")
	}
}

// Run with: go test -run '^TestTransfers$' -count=50 -timeout=10s
// The unchanged starter can hang because opposing transfers lock in opposite
// orders. After repair, every run must finish and conserve the total.
func TestTransfers(t *testing.T) {
	a, b := NewAccount(1, 1000), NewAccount(2, 1000)
	var wg sync.WaitGroup
	for worker := 0; worker < 16; worker++ {
		from, to := a, b
		if worker%2 == 1 {
			from, to = b, a
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 100; i++ {
				if !Transfer(from, to, 1) {
					t.Errorf("transfer failed")
				}
			}
		}()
	}
	wg.Wait()
	if got := a.Balance() + b.Balance(); got != 2000 {
		t.Fatalf("total = %d, want 2000", got)
	}
}
