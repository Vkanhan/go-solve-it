package main 

import(
	"fmt"
)

//Replace the each element in the array with the greatest element to its right.
func replaceWithGreatest(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return nums
	}

	result := make([]int, len(nums))
	maxRight := -1

	for i := n-1; i >= 0; i-- {
		result[i] = maxRight
		if nums[i] > maxRight {
			maxRight = nums[i]
		}
	}
	return result 
}

func main() {
	arr := []int{17, 18, 5, 4, 6, 1}

	fmt.Println(replaceWithGreatest(arr))
}