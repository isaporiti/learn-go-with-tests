package wallet

import "testing"

func TestWallet(t *testing.T) {

	assertBalance := func(t *testing.T, w *Wallet, want Bitcoin) {
		got := w.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(10)

		want := Bitcoin(10)
		assertBalance(t, &wallet, want)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		_ = wallet.Withdraw(10)

		want := Bitcoin(10)
		assertBalance(t, &wallet, want)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, &wallet, startingBalance)

		got := err.Error()
		want := "cannot withdraw, insufficient funds"
		if got != want {
			t.Errorf("want %q, got %q", want, got)
		}
	})
}
