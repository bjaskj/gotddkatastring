package strcalc

import (
	"testing"
)

func TestStringCalculator_EmptyString_ReturnsZero(t *testing.T) {
	var v = StringCalculator("")

	if v != 0 {
		t.Error("Expected 0, got ", v)
	}
}

func TestStringCalculator_OneNumber_ReturnsNumber(t *testing.T) {
	var v = StringCalculator("1")

	if v != 1 {
		t.Error("Expected 1, got ", v)
	}
}

func TestStringCalculator_TwoNumberSeparatedByDelimiter_ReturnsSum(t *testing.T) {
	var v = StringCalculator("1,2")

	if v != 3 {
		t.Error("Expected 3, got ", v)
	}
}
