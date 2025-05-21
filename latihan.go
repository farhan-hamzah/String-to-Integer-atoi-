package main

import (
	"fmt"
	"math"
	"unicode"
)

func myAtoi(s string) int {
	i := 0
	n := len(s)
	sign := 1
	result := 0

	// Skip leading spaces
	for i < n && s[i] == ' ' {
		i++
	}

	// Check sign
	if i < n {
		if s[i] == '-' {
			sign = -1
			i++
		} else if s[i] == '+' {
			i++
		}
	}

	// Read digits
	for i < n && unicode.IsDigit(rune(s[i])) {
		digit := int(s[i] - '0')

		// Check overflow
		if result > (math.MaxInt32-digit)/10 {
			if sign == 1 {
				return math.MaxInt32
			} else {
				return math.MinInt32
			}
		}

		result = result*10 + digit
		i++
	}

	return result * sign
}

func main() {
	fmt.Println(myAtoi("   -42"))             // -42
	fmt.Println(myAtoi("4193 with words"))    // 4193
	fmt.Println(myAtoi("words and 987"))      // 0
	fmt.Println(myAtoi("-91283472332"))       // -2147483648
}
