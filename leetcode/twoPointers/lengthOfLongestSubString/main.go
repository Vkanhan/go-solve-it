package main  

import(
	"fmt"
)

//lengthOfLongestSubstring function find the length of the longest substring without repeating characters
func lengthOfLongestSubstring(s string) int {
	 // Map to keep track of characters and their indices
	charMap := make(map[byte]int)

	//Initialize two pointers 
	left, maxLength := 0, 0

	for right := 0; right < len(s); right++ {
		if index, ok := charMap[s[right]]; ok && index >= left {
			left = index + 1	// Move left pointer if duplicate character is found
		}
		charMap[s[right]] = right	// Update the index of the character
		if currentLength := right - left + 1; currentLength > maxLength {
			maxLength = currentLength	// Update max length if needed
		}
	}
	// Return the maximum length
	return maxLength
}

func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))
}