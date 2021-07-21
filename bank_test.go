package money_test

import (
	"testing"

	"github.com/asahnoln/tdd-money"
)

func TestSimpleAddition(t *testing.T) {
	five := money.NewDollar(5)
	sum := five.Plus(five)
	bank := money.Bank{}
	reduced := bank.Reduce(sum, "USD")

	assertEquals(t, money.NewDollar(10), reduced)
}
