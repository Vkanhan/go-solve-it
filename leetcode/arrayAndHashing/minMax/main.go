package main 

import(
	"fmt"
)

func findMinMax(arr []int) (int, int) {
	min, max := arr[0], arr[0]

	for _, num := range arr {
		if num < min {
			min = num 
		}
		if num > max {
			max = num 
		}
	}
	return min, max 
}

func main() {
	arr := []int{3, 4, 5, 1, 8}
	fmt.Println(findMinMax(arr))
}