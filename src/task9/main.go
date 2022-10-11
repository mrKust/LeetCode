package main

import "fmt"

func main() {

	inputString := "(]"

	result := isValid(inputString)

	fmt.Println(result)
}

func isValid(s string) bool {

	queue := make([]uint8, 0)

	for i := 0; i < len(s); i++ {

		tmpChar := s[i]
		if tmpChar == ')' || tmpChar == '}' || tmpChar == ']' {
			if len(queue) == 0 {
				return false
			}
			previousChar := queue[len(queue)-1]
			if tmpChar == ')' {
				if previousChar == '(' {
					queue = queue[:len(queue)-1]
					continue
				} else {
					return false
				}
			}
			if tmpChar == '}' {
				if previousChar == '{' {
					queue = queue[:len(queue)-1]
					continue
				} else {
					return false
				}
			}
			if tmpChar == ']' {
				if previousChar == '[' {
					queue = queue[:len(queue)-1]
					continue
				} else {
					return false
				}
			}

		} else {
			queue = append(queue, tmpChar)

		}

	}

	if len(queue) == 0 {
		return true
	} else {
		return false
	}

}
