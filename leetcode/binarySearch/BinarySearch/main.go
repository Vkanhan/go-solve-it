package main 

import(
	"fmt"
)

func binarySearch(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	left, right := 0, len(nums) - 1

	for left < right {
		mid := left + (right - left) / 2

		switch {
		case nums[mid] < target:
			left = mid + 1
		case nums[mid] > target:
			right = mid - 1
		default:
			return mid 
		}
	}
	return -1
}

func main() {
    nums := []int{-1, 0, 3, 5, 9, 12}
    target := 9
	fmt.Println(binarySearch(nums, target))
}