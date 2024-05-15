package main

import (
	"errors"
	"fmt"
	"testing"
)

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

type Stringer interface {
	Stringer() string
}

func (w *Wallet) Desposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return errors.New("impossivel realizar withdraw, saldo insuficiente")
	}

	w.balance -= amount
	return nil
}

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Desposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	//ajudante no tratamento de erros
	assertError := func(t testing.TB, got error, want string) {
		t.Helper()

		if got == nil {
			t.Fatal("Nenhum erro encontrado")
		}

		if got.Error() != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	t.Run("withdraw saldo insuficiente", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, "impossivel realizar withdraw, saldo insuficiente")
		assertBalance(t, wallet, startingBalance)
	})

}
