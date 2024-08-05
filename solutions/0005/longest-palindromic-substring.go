package main

func longestPalindrome(s string) string {
	res := ""
	for left := 0; left < len(s); left++ {
		for right := left; right < len(s); right++ {
			substring := s[left : right+1]
			if isPalindrome(substring) {
				if len(substring) > len(res) {
					res = substring
				}
			}
		}
	}
	return res
}

func isPalindrome(s string) bool {
	left := 0
	right := len(s) - 1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
