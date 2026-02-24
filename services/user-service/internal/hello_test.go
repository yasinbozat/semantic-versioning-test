package internal

import "testing"

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	if result != 5 {
		t.Errorf("expected 5, got %d", result)
	}
}

func TestDiscount(t *testing.T) {
	result := Discount(100.0, 20.0)
	if result != 80.0 {
		t.Errorf("expected 80.0, got %f", result)
	}
}
