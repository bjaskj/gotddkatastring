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

func TestStringCalculator_MultipleNumbersSeparatedByDelimiter_ReturnsSum(t *testing.T) {
	testCases := map[int64]string{
		3: "1,2",
		10: "1,2,3,4",
	}

	for expected, input := range testCases {
		var v = StringCalculator(input)

		if v != expected {
			t.Error("Expected", expected, ", got ", v)
		}
	}
}

func TestStringCalculator_NewLineAsDelimiter_ReturnsSum(t *testing.T) {
	var v = StringCalculator("1\n2")

	if v != 3 {
		t.Error("Expected 3, got ", v)
	}
}
