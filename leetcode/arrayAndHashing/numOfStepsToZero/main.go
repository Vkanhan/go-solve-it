package main 

import(
	"fmt"
)

//numberOfSteps returns the number of steps to reduce a number to zero
func numberOfSteps(num int) int {
	steps := 0

	for num > 0 {
		switch {
		case num % 2 == 0:
			num /= 2 
		default:
			num -= 1
		}
		steps++
	}
	return steps
}

func main() {
	num := 4

	fmt.Println(numberOfSteps(num))
}