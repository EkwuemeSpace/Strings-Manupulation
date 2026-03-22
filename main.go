package main

import (
	"fmt"
	"strings"
)

/*
Write a program with a function called `analyzeString` that:
1. Takes any string as input
2. Prints total **character count**
3. Prints total **word count**
4. Prints the **reversed** string
5. Prints if it is a **palindrome**
6. Prints the **most frequent** character and its count
7. Prints the string with **duplicates removed**
*/

func countCharcter(s string) int {
	return len(strings.TrimSpace(s))
}

func wordCOunter(s string) int {
	return len(strings.Fields(strings.TrimSpace(s)))
}
func analyzeString(s string) {
	fmt.Println("Analising string please wait...")

	fmt.Println("character:", countCharcter(s))
	fmt.Println("word:", wordCOunter(s))
	fmt.Println("reverse", reveres(s))
	fmt.Println("palindrome", isPalindrome(s))
	// no idea on how to access most frequent character


	fmt.Println("No duplictae:", removeDuplicate(s))

}

func main() {
	analyzeString("racecar is a racecar")
}
