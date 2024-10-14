package main 

import(
	"fmt"
)

func lengthOfLastWord(s string) int {
	count := 0 

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' && count == 0 {
			continue
		}
		if s[i] != ' ' {
			count += 1
		} else {
			break
		}
	}
	return count 
}


func main() {
	fmt.Println(lengthOfLastWord("Einstein and Relativity  "))
}
