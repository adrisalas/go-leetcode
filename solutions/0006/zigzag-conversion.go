package main

import "strings"

func convert(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}

	rows := make([]string, numRows)
	rowIndex := 0
	right := true

	for i := range s {
		rows[rowIndex] += string(s[i])

		if right {
			rowIndex++
		} else {
			rowIndex--
		}

		switch rowIndex {
		case 0:
			right = true
		case numRows - 1:
			right = false
		}
	}
	return strings.Join(rows, "")
}
