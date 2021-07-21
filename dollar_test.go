package currency_test

import (
	"testing"

	"example.com/currency"
)

func TestMultiplication(t *testing.T) {
	five := currency.NewDollar(5);

	product := five.Times(2)
	assertEquals(t, 10, product.Amount)

	product = five.Times(3)
	assertEquals(t, 15, product.Amount)
}

func assertEquals(t testing.TB, want, got int) {
	t.Helper()
	if want != got {
		t.Errorf("wanted amount %v, got %v", want, got)
	}
}
