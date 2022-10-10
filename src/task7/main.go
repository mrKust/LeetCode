package main

import (
	"fmt"
)

func main() {

	var testArra = [...]int{2, 7, 11, 15}
	target := 9
	var resultArray = twoSum(testArra[:], target)
	fmt.Printf("%v", resultArray)

}

func twoSum(nums []int, target int) []int {

	for i := 0; i < len(nums) - 1; i++ {
		for k := i + 1; k < len(nums); k++ {
			if (nums[i] + nums[k]) == target {
				var resultArray = [2]int{i, k}
				return resultArray[:]
			}
		}
		fmt.Println(nums[i])
	}

	return nums
}
