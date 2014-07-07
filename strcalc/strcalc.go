package strcalc

import (
	"strconv"
)

func StringCalculator(input string) int64 {
	if input == "" {
		return 0
	}

	i, _ := strconv.ParseInt(input, 0, 32)

	return i
}
