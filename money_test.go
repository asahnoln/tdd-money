package money_test

import (
	"testing"

	"example.com/money"
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

func assertTrue(t testing.TB, got bool) {
	t.Helper()
	assertBool(t, true, got)
}

func assertFalse(t testing.TB, got bool) {
	t.Helper()
	assertBool(t, false, got)
}

func assertBool(t testing.TB, want, got bool) {
	t.Helper()
	if want != got {
		t.Errorf("want bool %v, got %v", want, got)
	}
}

func assertEquals(t testing.TB, want, got interface{}) {
	t.Helper()
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}
