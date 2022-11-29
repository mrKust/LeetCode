package main

import "fmt"

func main() {

	fmt.Println(isPalindrome(13231))
}

func isPalindrome(x int) bool {

	chars := []rune(fmt.Sprintf("%d", x))
	charsSize := len(chars)
	for i := 0; i < charsSize/2; i++ {
		charForward := chars[i]
		charReverse := chars[charsSize-1-i]
		fmt.Printf("%c - %c\n", charForward, charReverse)
		if charForward != charReverse {
			return false
		}
	}

	return true
}
