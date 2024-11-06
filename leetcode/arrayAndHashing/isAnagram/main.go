package main

import (
	"fmt"
	"sort"
)

// isAnagram checks if two strings are anagrams of each other
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	charCount := make(map[rune]int)

	for _, char := range s {
		charCount[char]++
	}

	for _, char := range t {
		charCount[char]--
	}

	for _, count := range charCount {
		if count != 0 {
			return false
		}
	}
	return true
}

func areAnagrams(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sArr := []rune(s)
	tArr := []rune(t)

	sort.Slice(sArr, func(i, j int) bool {
		return sArr[i] < sArr[j]
	})

	sort.Slice(tArr, func(i, j int) bool {
		return tArr[i] < tArr[j]
	})

	return string(sArr) == string(tArr)
}

func main() {

	s1, t1 := "anagram", "nagaram"
	fmt.Println(isAnagram(s1, t1))

	s2, t2 := "rat", "car"
	fmt.Println(isAnagram(s2, t2))

	s := "listen"
	t := "silent"
	fmt.Println(areAnagrams(s, t))


}
