package main

import (
	"fmt"
	"strconv"
)

// FizzBuzz function outputs fizz buzz according to the input
func FizzBuzz(num int) []string {
	var result []string

	for i := 1; i <= num; i++ {
		switch {
		case i%15 == 0:
			result = append(result, "FizzBuzz")
		case i%3 == 0:
			result = append(result, "Fizz")
		case i%5 == 0:
			result = append(result, "Buzz")
		default:
			result = append(result, strconv.Itoa(i))
		}
	}
	return result
}

func main() {
	num := 3

	fmt.Println(FizzBuzz(num))
}
