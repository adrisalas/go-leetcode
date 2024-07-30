package main

import "math"

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	size1 := len(nums1)
	size2 := len(nums2)

	if size1 > size2 {
		return findMedianSortedArrays(nums2, nums1)
	}

	oddSize := (size1+size2)%2 == 1
	half := (size1 + size2 + 1) / 2
	l := 0
	r := size1

	for l < r {
		p1 := l + (r-l)/2
		p2 := half - p1

		if nums1[p1] < nums2[p2-1] {
			l = p1 + 1
		} else {
			r = p1
		}
	}

	p1 := l
	p2 := half - p1

	n1 := math.MinInt
	if p1 > 0 && p2 > 0 {
		n1 = max(nums1[p1-1], nums2[p2-1])
	} else if p1 > 0 {
		n1 = nums1[p1-1]
	} else if p2 > 0 {
		n1 = nums2[p2-1]
	}

	if oddSize {
		return float64(n1)
	}

	n2 := math.MaxInt
	if p1 < size1 && p2 < size2 {
		n2 = min(nums1[p1], nums2[p2])
	} else if p1 < size1 {
		n2 = nums1[p1]
	} else if p2 < size2 {
		n2 = nums2[p2]
	}
	return float64(n1+n2) / 2
}
