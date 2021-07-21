package money_test

import "testing"

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
