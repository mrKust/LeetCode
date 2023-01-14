package main

import "fmt"

func main() {
	x := 5
	fmt.Println(mySqrt(x))
}

func mySqrt(x int) int {
	result := 1

	for result <= x {
		resInPow := result * result
		if resInPow == x {
			return result
		} else if resInPow > x {
			return result - 1
		} else {
			result++
		}

	}

	return 1
}
