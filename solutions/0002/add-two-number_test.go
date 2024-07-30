package main

import (
	"testing"
)

type TestCase struct {
	l1     *ListNode
	l2     *ListNode
	output *ListNode
}

func Test0002(t *testing.T) {
	testCases := []TestCase{
		{
			&ListNode{2, &ListNode{4, &ListNode{3, nil}}},
			&ListNode{5, &ListNode{6, &ListNode{4, nil}}},
			&ListNode{7, &ListNode{0, &ListNode{8, nil}}},
		},
		{
			&ListNode{0, nil},
			&ListNode{0, nil},
			&ListNode{0, nil},
		},
		{
			&ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}}}},
			&ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}},
			&ListNode{8, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{1, nil}}}}}}}},
		},
	}

	for _, tc := range testCases {
		res := addTwoNumbers(tc.l1, tc.l2)
		pivot := tc.output
		for pivot != nil {
			if pivot.Val != res.Val {
				t.Errorf("FAIL: l1 %v, l2 %v => expected %v, got %v", tc.l1, tc.l2, tc.output, res)
			}

			res = res.Next
			pivot = pivot.Next
		}
	}
}
