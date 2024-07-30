package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	root := &ListNode{}
	carry := 0
	currentNode := root

	for l1 != nil || l2 != nil || carry != 0 {
		currentNode.Next = &ListNode{carry, nil}
		currentNode = currentNode.Next

		if l1 != nil {
			currentNode.Val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			currentNode.Val += l2.Val
			l2 = l2.Next
		}

		carry = currentNode.Val / 10
		currentNode.Val = currentNode.Val % 10
	}

	return root.Next
}
