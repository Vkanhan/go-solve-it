package main 

import(
	"fmt"
)

// Finds the length of the longest substring with at most 'k' character replacements.
func characterReplacement(s string, k int) int {
    left := 0
    maxLength := 0
    maxCount := 0
    count := make(map[byte]int)

    for right := 0; right < len(s); right++ {
        count[s[right]]++
        
        // Update maxCount to reflect the highest frequency of any character in the window
        if count[s[right]] > maxCount {
            maxCount = count[s[right]]
        }

        // If the current window size minus the count of the most frequent character
        // is greater than k, shrink the window from the left
        if right-left+1-maxCount > k {
            count[s[left]]--
            left++
        }

        // Update the maxLength for the longest valid substring
        if right-left+1 > maxLength {
            maxLength = right - left + 1
        }
    }
    return maxLength
}


func main() {
    s1 := "XYYX"
    k1 := 2
    fmt.Println(characterReplacement(s1, k1)) 

    s2 := "AAABABB"
    k2 := 1
    fmt.Println(characterReplacement(s2, k2)) 
}

