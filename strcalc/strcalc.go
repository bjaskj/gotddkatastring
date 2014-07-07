package strcalc

import (
	"strconv"
	"strings"
)

const separator = ","

func StringCalculator(input string) int64 {
	if input == "" {
		return 0
	}

	values := strings.Split(input, separator)

	var sum int64

	for _, v := range values {
		i, _ := strconv.ParseInt(v, 0, 32)
		sum += i
	}

	return sum
}
