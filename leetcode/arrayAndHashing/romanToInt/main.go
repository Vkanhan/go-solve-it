package main 

import(
	"fmt"
)

func RomanToInt(s string) int {
	romanMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	result := 0
	greatest := 0

	//iterate over the string
	for i := len(s) - 1; i >= 0; i-- {
		letter := s[i]
		num := romanMap[string(letter)]

		if num >= greatest {
			greatest = num
			result += num
			continue
		}
		result -= num
	}
	return result
}

func main() {
	fmt.Println("Roman to Integer")
	fmt.Println(RomanToInt("CMXI"))
}