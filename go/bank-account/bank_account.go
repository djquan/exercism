package account

import "sync"

//Account represents a bank account
type Account struct {
	sync.RWMutex
	balance int64
	closed  bool
}

//Open creates an account with a given initial deposit
func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}

	return &Account{
		balance: initialDeposit,
		closed:  false,
	}
}

//Balance returns the balance of an account
func (a *Account) Balance() (int64, bool) {
	a.RLock()
	defer a.RUnlock()

	return a.balance, !a.closed
}

//Deposit adds money to the account
func (a *Account) Deposit(d int64) (int64, bool) {
	a.Lock()
	defer a.Unlock()

	b := a.balance + d
	if b < 0 {
		return a.balance, false
	}

	if a.closed {
		return 0, false
	}

	a.balance = b
	return a.balance, true
}

//Close adds money to the account
func (a *Account) Close() (int64, bool) {
	a.Lock()
	defer a.Unlock()

	if a.closed {
		return 0, false
	}

	a.closed = true
	payout := a.balance
	a.balance = 0

	return payout, true
}
