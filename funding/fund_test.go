package funding

import (
	"sync"
	"testing"
)

const WORKERS = 10

func BenchmarkFund(b *testing.B) {
	if b.N < WORKERS {
		return
	}

	fund := NewFund(b.N)
	dollarsPerFounder := b.N / WORKERS
	var wg sync.WaitGroup

	for i := 0; i < WORKERS; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < dollarsPerFounder; i++ {
				fund.Withdraw(1)
			}
		}()
	}

	wg.Wait()

	if fund.Balance() != 0 {
		b.Errorf("Balance was not 0 : %d", fund.Balance())
	}
}

func BenchmarkWithdrawals(b *testing.B) {
	if b.N < WORKERS {
		return
	}

	server := NewFundServer(b.N)
	dollarsPerFounder := b.N / WORKERS
	var wg sync.WaitGroup
	for i := 0; i < WORKERS; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			pizzaTime := false
			for i := 0; i < dollarsPerFounder; i++ {
				server.Transact(func(f *Fund) {
					if f.Balance() <= 10 {
						pizzaTime = true
						return
					}
					f.Withdraw(1)
				})
				if pizzaTime {
					break
				}
			}
		}()
	}

	wg.Wait()

	balance := server.Balance()

	if balance != 10 {
		b.Error("Balance wasnt ten $: ", balance)
	}
}
