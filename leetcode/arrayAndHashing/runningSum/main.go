package main

import(
	"fmt"
)

//runningSum calculates the running sum of the 1D array
func runningSum(nums []int) []int {
	result := make([]int, len(nums))
	sum := 0

	for i, num := range nums {
		sum += num
		result[i] = sum
	}
	return result
}

func main() {
	nums := []int{2, 4, 6, 3, 7}

	fmt.Println(runningSum(nums))
}