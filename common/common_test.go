package common

import (
	"testing"
)

func TestBidFees(t *testing.T) {
	o := &CommonOrder{"cme", "btc", 5, 10.0, 0.01, Bid}

	if GetFeeAdjustedPrice(o) != o.price*(1-Price(o.fee)) {
		t.Error()
	}
}

func TestAskFees(t *testing.T) {
	o := &CommonOrder{"cme", "btc", 5, 10.0, 0.01, Ask}

	if GetFeeAdjustedPrice(o) != o.price*(1+Price(o.fee)) {
		t.Error()
	}
}

func TestNoFees(t *testing.T) {
	o := &CommonOrder{"cme", "btc", 5, 10.0, 0.0, Bid}

	if GetFeeAdjustedPrice(o) != o.price {
		t.Error()
	}
}

func TestNegativeBidFees(t *testing.T) {
	o := &CommonOrder{"cme", "btc", 5, 10.0, -0.01, Bid}

	if GetFeeAdjustedPrice(o) != o.price*(1-Price(o.fee)) {
		t.Error()
	}
}

func TestNegativeAskFees(t *testing.T) {
	o := &CommonOrder{"cme", "btc", 5, 10.0, -0.01, Ask}

	if GetFeeAdjustedPrice(o) != o.price*(1+Price(o.fee)) {
		t.Error()
	}
}
