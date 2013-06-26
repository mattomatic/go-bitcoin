package common

import (
	"testing"
)

func TestBudget(t *testing.T) {
	eb := NewExchangeBudget()

	b := eb.Get("CME")

	if len(eb) != 1 {
		t.Error()
	}

	if b.USD != 0 {
		t.Error()
	}

	if b.BTC != 0 {
		t.Error()
	}

	b.USD = 1
	b.BTC = 2

	b2 := eb.Get("CME")

	if b2.USD != 1 {
		t.Error()
	}

	if b2.BTC != 2 {
		t.Error()
	}
}
