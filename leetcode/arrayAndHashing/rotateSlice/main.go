package main  

import(
	"fmt"
)

//rotateSlice rotates the slice to the right by given number of steps
func rotateSlice(nums []int, index int) []int {
	result := []int{}

	// Append elements from nums starting from index + 1 to the end of the slice.
	for i := index + 1; i < len(nums); i++ {
		result = append(result, nums[i])
	}

	// Append the remaining elements from the beginning of nums to index.
	for i := 0; i <= index; i++ {
		result = append(result, nums[i])
	}
	return result
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	index := 2

	fmt.Println(rotateSlice(nums, index))
}
