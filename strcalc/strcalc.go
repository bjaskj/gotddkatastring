package strcalc

import (
	"strconv"
	"strings"
)

const separatorComma = ","
const separatorNewLine = "\n"

func StringCalculator(input string) int64 {
	if input == "" {
		return 0
	}

	var values = []string{}
	valuesFromComma := strings.Split(input, separatorComma)

	for _, v := range valuesFromComma {
		valuesFromNewLine := strings.Split(v, separatorNewLine)

		values = append(values, valuesFromNewLine...)
	}

	var sum int64

	for _, v := range values {
		i, _ := strconv.ParseInt(v, 0, 32)
		sum += i
	}

	return sum
}
