package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

	tree3L1 := TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}

	tree3L2 := TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}

	tree2L1 := TreeNode{
		Val:   2,
		Left:  &tree3L1,
		Right: &tree3L2,
	}

	tree3R1 := TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}

	tree3R2 := TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}

	tree2R1 := TreeNode{
		Val:   2,
		Left:  &tree3R1,
		Right: &tree3R2,
	}

	root := TreeNode{
		Val:   1,
		Left:  &tree2L1,
		Right: &tree2R1,
	}

	fmt.Print(isSymmetric(&root))
}

func isSymmetric(root *TreeNode) bool {

	return helper(root.Left, root.Right)
}

func helper(left *TreeNode, right *TreeNode) bool {

	if (left == nil) && (right == nil) {
		return true
	}
	if ((left == nil) && (right != nil)) || ((left != nil) && (right == nil)) {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return helper(left.Left, right.Right) && helper(left.Right, right.Left)
}
