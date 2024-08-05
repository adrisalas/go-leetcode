package main

import "strconv"

func countSeniors(details []string) int {
	count := 0

	for _, v := range details {
		age, _ := strconv.Atoi(v[11:13])
		if age > 60 {
			count++
		}
	}
	return count
}
