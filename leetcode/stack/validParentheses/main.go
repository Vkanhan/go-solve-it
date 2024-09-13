package main

import (
	"fmt"
)

// Check if the parentheses is Valid or not
func isValidParentheses(s string) bool {
	if len(s) == 0 || len(s)%2 == 1 {
		return false
	}
	stack := []rune{}
	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, char := range s {
		if matchingChar, ok := pairs[char]; ok {
			// If the stack is empty or the top of the stack is not the matching opening bracket
			if len(stack) == 0 || stack[len(stack)-1] != matchingChar {
				return false
			}
			// Pop the matching opening bracket
			stack = stack[:len(stack)-1]
		} else {
			// Push the opening bracket onto the stack
			stack = append(stack, char)
		}
	}

	// Check if the stack is empty
	return len(stack) == 0
}

func main() {
	examples := []string{"()", "()[]{}", "(]", "([)]", "{[]}"}
	for _, example := range examples {
		fmt.Printf("Is \"%s\" valid? %v\n", example, isValidParentheses(example))
	}
}
