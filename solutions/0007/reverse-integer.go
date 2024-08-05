package main

import "math"

func reverse(x int) int {
	neg := false

	if x < 0 {
		neg = true
		x = x * -1
	}

	o := 0

	for x > 0 {
		o = (o * 10) + (x % 10)
		x = x / 10
	}

	if o > math.MaxInt32 {
		return 0
	}

	if neg {
		return o * -1
	}

	return o
}
