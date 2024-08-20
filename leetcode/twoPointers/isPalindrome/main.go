package main

import (
	"fmt"
	"unicode"
)

//isPalindrome checks if the string is a valid palindrome
func isPalindrome(s string) bool {
	var cleaned []rune

	//Clean the string
	for _, char := range s {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			cleaned = append(cleaned, unicode.ToLower(char))
		}
	}

	//iterate through the slice checking if the first and last elements are the same
	length := len(cleaned)
	for i := 0; i < length/2; i++ {
		if cleaned[i] != cleaned[length-1-i] {
			return false
		}
	}

	return true
}

//IsPalindromeWord checks if the word is palindrome or not
func isPalindromeWord(s string) bool {
    left := 0
    right := len(s) - 1

    // Loop until the two pointers meet in the middle
    for left < right {
        if s[left] != s[right] {
            return false
        }
        left++
        right--
    }
    return true
}

func main() {
	s1 := "Was it a car or a cat I saw?"
	fmt.Println(isPalindrome(s1)) 

	s2 := "tab a cat"
	fmt.Println(isPalindrome(s2)) 

    s3 := "racecar"
    fmt.Println(isPalindromeWord(s3))
}
