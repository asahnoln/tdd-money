package money_test

import (
	"testing"

	"github.com/asahnoln/tdd-money"
)

func TestMultiplication(t *testing.T) {
	five := money.NewDollar(5)
	assertEquals(t, money.NewDollar(10), five.Times(2))
	assertEquals(t, money.NewDollar(15), five.Times(3))
}

func TestEquality(t *testing.T) {
	assertTrue(t, money.NewDollar(5).Equals(money.NewDollar(5)))
	assertFalse(t, money.NewDollar(5).Equals(money.NewDollar(6)))
	assertFalse(t, money.NewFranc(5).Equals(money.NewDollar(5)))
}

func TestCurrency(t *testing.T) {
	assertEquals(t, "USD", money.NewDollar(1).Currency())
	assertEquals(t, "CHF", money.NewFranc(1).Currency())
}

func TestPlusReturnsSum(t *testing.T) {
	five := money.NewDollar(5)
	sum := five.Plus(five)
	assertEquals(t, five, sum.Augend)
	assertEquals(t, five, sum.Addend)
}

func TestReduceMoney(t *testing.T) {
	bank := money.NewBank()
	got := bank.Reduce(money.NewDollar(1), "USD")
	assertEquals(t, money.NewDollar(1), got)
}

func TestSimpleAddition(t *testing.T) {
	five := money.NewDollar(5)
	sum := five.Plus(five)
	bank := money.NewBank()
	reduced := bank.Reduce(sum, "USD")

	assertEquals(t, money.NewDollar(10), reduced)
}

func TestReduceMoneyDifferentCurrency(t *testing.T) {
	bank := money.NewBank()
	bank.AddRate("CHF", "USD", 2)
	got := bank.Reduce(money.NewFranc(2), "USD")

	assertEquals(t, money.NewDollar(1), got)
}

func TestMixedAddition(t *testing.T) {
	fiveBucks := money.NewDollar(5)
	tenFrancs := money.NewFranc(10)

	bank := money.NewBank()
	bank.AddRate("CHF", "USD", 2)
	got := bank.Reduce(fiveBucks.Plus(tenFrancs), "USD")

	assertEquals(t, money.NewDollar(10), got)
}
