package main

func main() {

	inorderTraversal()
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {

	var resultArray []int

	helper(root, &resultArray)

	return resultArray
}

func helper(root *TreeNode, resultArray *[]int) {
	if root != nil {
		helper(root.Left, resultArray)
		*resultArray = append(*resultArray, root.Val)
		helper(root.Right, resultArray)
	}
}
