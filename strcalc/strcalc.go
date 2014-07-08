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
	if strings.HasPrefix(input, "//") {
		// we have prefix to change delimiter
		delimiter := string(input[2])
		values = strings.Split(input[4:], delimiter)
	} else {
		// we don't have prefix to change delimiter
		valuesFromComma := strings.Split(input, separatorComma)

		for _, v := range valuesFromComma {
			valuesFromNewLine := strings.Split(v, separatorNewLine)

			values = append(values, valuesFromNewLine...)
		}
	}

	var sum int64

	for _, v := range values {
		i, _ := strconv.ParseInt(v, 0, 32)
		sum += i
	}

	return sum
}
