// 21. Merge Two Sorted Lists

package main

import (
	"fmt"
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func val(node *ListNode) int {
	if node == nil {
		return math.MaxInt64
	}
	return node.Val
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	stab := &ListNode{}
	current := stab

	for list1 != nil || list2 != nil {
		if val(list1) < val(list2) {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	return stab.Next
}

func main() {
	node_a := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{Val: 4},
		},
	}

	node_b := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  3,
			Next: &ListNode{Val: 4},
		},
	}
	node_res := mergeTwoLists(node_a, node_b)

	for node_res.Next != nil {
		fmt.Println(node_res.Val)
		node_res = node_res.Next
	}
}
