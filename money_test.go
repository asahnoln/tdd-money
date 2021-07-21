package currency_test

import (
	"testing"

	"example.com/currency"
)

func TestMultiplication(t *testing.T) {
	five := currency.NewDollar(5);
	assertEquals(t, currency.NewDollar(10), five.Times(2))
	assertEquals(t, currency.NewDollar(15), five.Times(3))
}

func TestFrancMultiplication(t *testing.T) {
	five := currency.NewFranc(5);
	assertEquals(t, currency.NewFranc(10), five.Times(2))
	assertEquals(t, currency.NewFranc(15), five.Times(3))
}

func TestEquality(t *testing.T) {
	assertTrue(t, currency.NewDollar(5).Equals(currency.NewDollar(5)))
	assertFalse(t, currency.NewDollar(5).Equals(currency.NewDollar(6)))
	assertTrue(t, currency.NewFranc(5).Equals(currency.NewFranc(5)))
	assertFalse(t, currency.NewFranc(5).Equals(currency.NewFranc(6)))
	assertFalse(t, currency.NewFranc(5).Equals(currency.NewDollar(5)))
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
		t.Errorf("want amount %v, got %v", want, got)
	}
}
