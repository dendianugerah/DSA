package main

import "fmt"

type ListNode struct {
	val  int
	next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var res *ListNode

	for head != nil {
		prev := head.next // 2
		head.next = res   // 1
		res = head        // 1
		head = prev       // 2
	}

	return res
}

func main() {
	head := &ListNode{val: 1}
	head.next = &ListNode{val: 2}

	head = reverseList(head)

	for head != nil {
		fmt.Println(head.val)
		head = head.next
	}
}
