package main

import "fmt"

func frequentChar(s string) (rune, int) {
	mostFreq := make(map[rune]int)

	for _, ch := range s {
		mostFreq[ch]++
	}

	var topChar rune
	topCount := 0

	for char, count := range mostFreq {
		if count > topCount {
			topCount = count
			topChar = char
		}
	}

	return topChar, topCount
}

func main() {
	fmt.Println(frequentChar("innocent"))
}
