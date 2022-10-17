package main

import "fmt"

func main() {

	var sortedArray = []int{1}
	//var sortedArray = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	resultK := removeDuplicates(sortedArray)
	fmt.Println(resultK)

}

func removeDuplicates(nums []int) int {

	result := 0

	for {
		if result >= len(nums) {
			break
		}
		if nums[result] == -1000 {
			break
		}
		if result+1 < len(nums) {
			for {
				if nums[result] == nums[result+1] {
					moveArrayToOnePoseLeft(&nums, result)
				} else {
					break
				}
			}
		}
		result++
	}

	return result
}

func moveArrayToOnePoseLeft(nums *[]int, pos int) {
	for i := pos; i < len(*nums)-1; i++ {
		(*nums)[i] = (*nums)[i+1]
	}
	(*nums)[len(*nums)-1] = -1000
}
