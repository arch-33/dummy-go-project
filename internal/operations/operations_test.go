package operations

import (
	"testing"
)

func TestSum(t *testing.T) {
	result := Sum(2, 3)
	expected := 6

	if result != expected {
		t.Errorf("Sum(2, 3) = %d; expected %d", result, expected)
	}
}

func TestSubtract(t *testing.T) {
	result := Subtract(5, 3)
	expected := 2

	if result != expected {
		t.Errorf("Subtract(5, 3) = %d; expected %d", result, expected)
	}
}

func TestMultiply(t *testing.T) {
	result := Multiply(4, 5)
	expected := 20

	if result != expected {
		t.Errorf("Multiply(4, 5) = %d; expected %d", result, expected)
	}
}

func TestDivide(t *testing.T) {
	result := Divide(10, 2)
	expected := 50

	if result != expected {
		t.Errorf("Divide(10, 2) = %d; expected %d", result, expected)
	}
}
