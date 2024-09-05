package main 

import(
	"fmt"
)

func reverseArray(nums []int) []int {
	n := len(nums)

    // Loop until the middle of the array (n/2), as we only need to swap half
	for i := 0; i < n/2; i++ {
		// Swap the element at index i with the corresponding element from the end (n-i-1)
		nums[i], nums[n-1-i] = nums[n-1-i], nums[i] 
	}
	return nums
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(reverseArray(nums))
}