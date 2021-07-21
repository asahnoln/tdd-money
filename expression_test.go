package money_test

import (
	"testing"

	"example.com/money"
)

func TestReduceSum(t *testing.T) {
	sum := money.Sum{money.NewDollar(3), money.NewDollar(4)}
	bank := money.Bank{}
	got := bank.Reduce(sum, "USD")
	assertEquals(t, money.NewDollar(7), got)
}
