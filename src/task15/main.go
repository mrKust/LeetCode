package main

import "fmt"

func main() {

	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2))

}

func searchInsert(nums []int, target int) int {

	leftIndex := 0
	rightIndex := len(nums) - 1

	for leftIndex <= rightIndex {

		searchIndex := (leftIndex + rightIndex) / 2

		if nums[searchIndex] == target {
			return searchIndex
		} else if target > nums[searchIndex] {
			leftIndex = searchIndex + 1
		} else {
			rightIndex = rightIndex - 1
		}

	}

	return rightIndex + 1
}
