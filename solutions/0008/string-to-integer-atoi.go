package main

import "strings"

func myAtoi(s string) int {
	s = strings.TrimSpace(s)

	if len(s) == 0 {
		return 0
	}

	negative := false

	if s[0] == '-' {
		negative = true
		s = s[1:]
	} else if s[0] == '+' {
		s = s[1:]
	}

	output := 0

	for i, char := range s {
		if char < '0' || char > '9' {
			break
		}

		output = output*10 + int(s[i]-'0')

		if output > 2147483647 && !negative {
			return 2147483647
		}

		if output > 2147483648 && negative {
			return -2147483648
		}
	}

	if negative {
		output = -output
	}

	return output
}
