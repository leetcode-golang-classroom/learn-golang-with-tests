package pointers_and_errors

import "fmt"

type Bitcoin int

// String - stringer for Bitcoin
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

// Deposit - deposit amount into wallet balance
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Balance - show bitcoin balance
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// Withdraw - deposit amount from wallet balance
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if w.balance < amount {
		return fmt.Errorf("oh no")
	}
	w.balance -= amount
	return nil
}
