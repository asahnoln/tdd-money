package money_test

import (
	"testing"

	"github.com/asahnoln/tdd-money"
)

func TestIdentityRate(t *testing.T) {
	assertEquals(t, 1, money.NewBank().Rate("USD", "USD"))
}
