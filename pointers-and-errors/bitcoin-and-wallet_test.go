package pointers_and_errors

import "testing"

// Bitcoin Wallet
func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	assertError := func(t testing.TB, err, want error) {
		t.Helper()
		if err == nil {
			t.Error("wanted an error but didn't get one")
		}

		if err != want {
			t.Errorf("got %q want %q", err.Error(), want)
		}
	}

	assertNoError := func(t testing.TB, got error) {
		t.Helper()
		if got != nil {
			t.Fatal("got an error but didn't want one")
		}
	}
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(10)
		want := Bitcoin(10)
		assertBalance(t, wallet, want)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		err := wallet.Withdraw(Bitcoin(10))
		want := Bitcoin(10)
		assertNoError(t, err)
		assertBalance(t, wallet, want)
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)

	})
}
