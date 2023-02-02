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

	fmt.Print(isBalanced(&root))
}

func isBalanced(root *TreeNode) bool {
	res, _ := dfs(root)
	return res
}

func dfs(root *TreeNode) (bool, int) {
	if root == nil {
		return true, 0
	}

	isLeftBalanced, leftHeight := dfs(root.Left)
	isRightBalanced, rightHeight := dfs(root.Right)
	diff := abs(leftHeight - rightHeight)
	if isLeftBalanced && isRightBalanced && (diff <= 1) {
		return true, 1 + max(leftHeight, rightHeight)
	}
	return false, -1
}

func abs(n int) int {
	if n < 0 {
		return -n
	} else {
		return n
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
