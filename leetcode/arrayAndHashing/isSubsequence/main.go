package main  

import(
	"fmt"
)

// isSubsequence checks if string s is a subsequence of string t
func isSubsequence(s , t string) bool {
	sLen, tLen := len(s), len(t)

	if sLen == 0 {
		return true //An empty string is always a subsequence of any string
	}

	sIndex, tIndex := 0, 0

	for tIndex < tLen {
		if s[sIndex] == t[tIndex] {
			sIndex++
		}
		if sIndex == sLen {
			return true 	//All characters of s are found in t
		}
		tIndex++
	}
	return false
}

func main() {
	s1 := "abc"
	t1 := "ahbgdc"
	fmt.Println(isSubsequence(s1, t1)) 

	s2 := "axc"
	t2 := "ahbgdc"
	fmt.Println(isSubsequence(s2, t2))
}