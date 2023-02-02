package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	tree31 := TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}

	tree32 := TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}

	tree21 := TreeNode{
		Val:   2,
		Left:  &tree31,
		Right: nil,
	}

	tree22 := TreeNode{
		Val:   3,
		Left:  nil,
		Right: &tree32,
	}

	root := TreeNode{
		Val:   1,
		Left:  &tree21,
		Right: &tree22,
	}

	fmt.Print(maxDepth(&root))
}

func maxDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}

	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func max(x int, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
