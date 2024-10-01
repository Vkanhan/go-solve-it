package main 

import(
	"fmt"
)

//bubble sort sorts in ascending order by comparing the adjacent inputs
func bubbleSort(input []int) []int {
	swapped := true 

	for swapped {
		swapped = false  

		for i := 1; i < len(input); i++ {
			if input[i] < input[i-1] {
				input[i], input[i-1] = input[i-1], input[i]
				swapped = true 
			}
		}
	}
	return input 
}

func main() {
	unSorted := []int{3, 1, 88, 42, 5}
	fmt.Println(bubbleSort(unSorted))
}