package mathutil

import "testing"

func TestAdd(t *testing.T) {
	result := Add(2, 3)

	if result != 5 {
		t.Errorf("Add(2,3) = %d; want 5", result)
	}
}

func TestMultiply(t *testing.T) {
	result := Multiply(2, 3)
	if result != 6 {
		t.Errorf("Multiply(2, 3) = %d; want 6", result)
	}
}
