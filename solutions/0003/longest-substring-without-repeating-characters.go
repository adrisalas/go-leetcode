package main

func lengthOfLongestSubstring(s string) int {
	output := 0
	count := make([]int, 256)

	left := 0
	for right := 0; right < len(s); right++ {
		rightChar := s[right]
		count[rightChar]++
		for count[rightChar] > 1 {
			leftChar := s[left]
			count[leftChar]--
			left++
		}
		windowSize := right - left + 1
		if windowSize > output {
			output = windowSize
		}
	}
	return output
}
