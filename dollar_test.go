package currency_test

import (
	"testing"

	"example.com/currency"
)

func TestMultiplication(t *testing.T) {
	five := currency.NewDollar(5);
	five.Times(2)

	want := 10
	got := five.Amount

	if want != got {
		t.Errorf("wanted amount %v, got %v", want, got)
	}
}
