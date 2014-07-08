package strcalc

import (
	"strconv"
	"strings"
)

const separatorComma = ","
const separatorNewLine = "\n"

func StringCalculator(input string) int64 {
	// return zero if no values
	if input == "" {
		return 0
	}

	var values = []string{}

	// check if we have prefix
	if hasDelimiter(input) {
		// we have prefix to change delimiter
		values = valuesByCustomDelimiter(input)
	} else {
		values = valuesByStandardDelimiter(input)
	}

	var sum int64

	for _, v := range values {
		i, _ := strconv.ParseInt(v, 0, 32)
		sum += i
	}

	return sum
}

func valuesByStandardDelimiter(input string) []string {
	var values = []string{}

	// we don't have prefix to change delimiter
	valuesFromComma := strings.Split(input, separatorComma)

	for _, v := range valuesFromComma {
		valuesFromNewLine := strings.Split(v, separatorNewLine)

		values = append(values, valuesFromNewLine...)
	}

	return values
}

func valuesByCustomDelimiter(input string) []string {
	delimiter := string(input[2])
	return strings.Split(input[4:], delimiter)
}

func hasDelimiter(input string) bool {
	return strings.HasPrefix(input, "//")
}
