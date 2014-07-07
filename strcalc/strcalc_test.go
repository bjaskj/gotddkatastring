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
