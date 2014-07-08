package strcalc

import (
	"testing"
)

func TestStringCalculator_EmptyString_ReturnsZero(t *testing.T) {
	var v, _ = StringCalculator("")

	if v != 0 {
		t.Error("Expected 0, got ", v)
	}
}

func TestStringCalculator_OneNumber_ReturnsNumber(t *testing.T) {
	var v, _ = StringCalculator("1")

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
		var v, _ = StringCalculator(input)

		if v != expected {
			t.Error("Expected", expected, ", got ", v)
		}
	}
}

func TestStringCalculator_NewLineAsDelimiter_ReturnsSum(t *testing.T) {
	var v, _ = StringCalculator("1\n2")

	if v != 3 {
		t.Error("Expected 3, got ", v)
	}
}

func TestStringCalculator_ChangeDelimiter_ReturnsSum(t *testing.T) {
	var v, _ = StringCalculator("//;\n1;2")

	if v != 3 {
		t.Error("Expected 3, got ", v)
	}
}

func TestStringCalculator_LessThanZero_ThrowsError(t *testing.T) {
	var _, err = StringCalculator("-1")

	if err != true {
		t.Error("Expected true, got ", err)
	}
}
