package main

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i := range nums {
		n2 := nums[i]
		n1 := target - n2
		p, ok := m[n1]
		if ok {
			return []int{p, i}
		}
		m[n2] = i
	}
	return []int{-1, -1}
}
