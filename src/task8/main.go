package main

import (
	"fmt"
)

func main() {

	var testArra = [2]string{"ab", "a"}
	var resultString = longestCommonPrefix(testArra[:])
	fmt.Printf("%s", resultString)

}

func longestCommonPrefix(strs []string) string {

	endCommonIndex := 0
	complite := false

	minLenght := len(strs[0])

	for i := 1; i < len(strs); i++ {
		if len(strs[i]) < minLenght {
			minLenght = len(strs[i])
		}
	}

	for k := 0; k < minLenght; k++ {
		currentElement := strs[0][k]
		for i := 1; i < len(strs); i++ {
			checkedElement := strs[i][k]
			if currentElement != checkedElement {
				complite = true
				break
			}

		}
		if complite {
			break
		} else {
			endCommonIndex++
		}
	}

	return strs[0][:endCommonIndex]
}
