package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	val5 := ListNode{
		Val:  3,
		Next: nil,
	}

	val4 := ListNode{
		Val:  3,
		Next: &val5,
	}

	val3 := ListNode{
		Val:  1,
		Next: &val4,
	}

	val2 := ListNode{
		Val:  1,
		Next: &val3,
	}

	val1 := ListNode{
		Val:  1,
		Next: &val2,
	}

	resultHead := deleteDuplicates(&val1)
	fmt.Println(resultHead)
}

func deleteDuplicates(head *ListNode) *ListNode {

	currentNode := head

	for (currentNode != nil) && (currentNode.Next != nil) {
		nextNode := currentNode.Next
		for (currentNode.Next != nil) && (nextNode.Val == currentNode.Val) {
			currentNode.Next = nextNode.Next
			nextNode = currentNode.Next
		}
		currentNode = currentNode.Next
	}

	return head
}
