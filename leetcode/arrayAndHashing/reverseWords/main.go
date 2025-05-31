package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	words := strings.Fields(s)

	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}

func reverse(s string) string {
	//extract the words manually
	words := [][]rune{}
	current := []rune{}

	for _, ch := range s {
		if ch != ' ' {
			current = append(current, ch)
		} else {
			if len(current) > 0 {
				words = append(words, current)
				current = []rune{}
			}
		}
	}
	//append the last word
	if len(current) > 0 {
		words = append(words, current)
	}

	//reverse the order of words
	result := []rune{}
	for i := len(words) - 1; i >= 0; i-- {
		result = append(result, words[i]...)
		if i != 0 {
			result = append(result, ' ')
		}
	}
	return string(result)
}

func main() {
	s := "Hello World 123"

	fmt.Println(reverseWords(s))

	fmt.Println(reverse(s))
}
