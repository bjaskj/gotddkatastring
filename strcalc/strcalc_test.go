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
		t.Error("Expected 0, got ", v)
	}
}
