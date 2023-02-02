package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

	nums := []int{0, 1, 2, 3, 4, 5, 6, 7}

	//nums := []int{1, 3}

	root := sortedArrayToBST(nums)
	fmt.Print(root)
}

func sortedArrayToBST(nums []int) *TreeNode {
	return helper(nums)
}

func helper(nums []int) *TreeNode {

	if len(nums) == 0 {
		return nil
	}

	rootValIndex := len(nums) / 2
	rootVal := nums[rootValIndex]

	root := TreeNode{
		Val: rootVal,
	}
	root.Left = helper(nums[0:rootValIndex])
	root.Right = helper(nums[rootValIndex+1:])

	return &root
}
