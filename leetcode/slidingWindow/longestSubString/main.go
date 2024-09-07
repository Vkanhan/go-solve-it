package main  

import(
	"fmt"
)

//longest substring without any duplicates
func longestSubString(s string) int {
	if len(s) == 0 {
		return 0
	}

	charSet := make(map[byte]bool)

	left, maxLen := 0, 0

	for right := 0; right < len(s); right++ {
		if charSet[s[right]] {
			delete(charSet, s[left])
			left++
		}

		charSet[s[right]] = true

		//right-left gives the length of substring and adding 1 because of 0 index
		if right - left + 1 > maxLen {
			maxLen = right - left + 1
		}
	}
	return maxLen
}

func main() {
	s := "abcdabcd"
	fmt.Println(longestSubString(s))
}