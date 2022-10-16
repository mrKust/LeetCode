package main

import (
	"fmt"
)

func main() {

	/*node23 := ListNode{Val: 4, Next: nil}
	node22 := ListNode{Val: 3, Next: &node23}
	node21 := ListNode{Val: 1, Next: &node22}

	node13 := ListNode{Val: 4, Next: nil}
	node12 := ListNode{Val: 2, Next: &node13}
	node11 := ListNode{Val: 1, Next: &node12}*/

	node21 := ListNode{Val: 0, Next: nil}

	resultNode := mergeTwoLists(nil, &node21)

	fmt.Print("\n", resultNode.Val)

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	var resultHead *ListNode
	i := -1
	var previousNode *ListNode
	var addedNode *ListNode
	for {

		i++

		if (list1 == nil) && (list2 == nil) {
			break
		}

		if i == 0 {
			if list1 == nil {
				addedNode = list2
				list2 = list2.Next
				resultHead = addedNode
				continue
			} else if list2 == nil {
				addedNode = list1
				list1 = list1.Next
				resultHead = addedNode
				continue
			}
			if list1.Val <= list2.Val {
				addedNode = list1
				list1 = list1.Next
				resultHead = addedNode

			} else {
				addedNode = list2
				list2 = list2.Next
				resultHead = addedNode
			}
		} else {
			if list2 == nil {
				previousNode = addedNode
				addedNode = list1
				previousNode.Next = addedNode
				list1 = list1.Next
				continue
			} else if list1 == nil {
				previousNode = addedNode
				addedNode = list2
				previousNode.Next = addedNode
				list2 = list2.Next
				continue
			}
			if list1.Val <= list2.Val {
				previousNode = addedNode
				addedNode = list1
				previousNode.Next = addedNode
				list1 = list1.Next

			} else {
				previousNode = addedNode
				addedNode = list2
				previousNode.Next = addedNode
				list2 = list2.Next
			}
		}

	}

	return resultHead
}
