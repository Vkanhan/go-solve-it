package main 

import(
	"fmt"
)

//reverseString outputs the reverse of the string
func reverseString(s string) string {
	// Convert string to rune slice to handle Unicode characters
	arr := []rune(s)
	// Initialize two pointers
	left, right := 0, len(arr) - 1

	for left < right {
		// Swap characters at left and right pointers
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
	return string(arr)
}

func main() {
	s := "hello"
    fmt.Println(reverseString(s))
}
