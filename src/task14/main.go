package main

import "fmt"

func main() {

	fmt.Println(removeElement([]int{3, 2, 2, 3}, 3))
}

func removeElement(nums []int, val int) int {

	returnIndex := 0
	getIndex := 0
	swapIndex := 0

	for i := 0; i < len(nums); i++ {
		if nums[getIndex] == val {
			nums[getIndex] = -1
			tmp := nums[len(nums)-1-swapIndex]
			nums[len(nums)-1-swapIndex] = nums[getIndex]
			nums[getIndex] = tmp
			swapIndex++
			continue
		}
		getIndex++
		returnIndex++
	}

	return returnIndex
}
