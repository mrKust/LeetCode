package main

import "fmt"

func main() {

	var digits = []int{1, 2, 3}

	resultArray := plusOne(digits)

	fmt.Println(resultArray)

}

func plusOne(digits []int) []int {

	for i := len(digits) - 1; i >= 0; i-- {
		fmt.Println(digits[i])
	}

	return digits
}
