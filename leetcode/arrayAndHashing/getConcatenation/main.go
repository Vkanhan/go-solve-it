package main

import (
	"fmt"
)

// concatenate nums with itself 2 times
func getConcatenation(nums []int) []int {
	ans := make([]int, 2*len(nums))

	for i := 0; i < len(nums); i++ {
		ans[i] = nums[i]
		ans[i+len(nums)] = nums[i]
	}
	return ans
}

//By using copy method
func getConcatenationByCopying(nums []int) []int {
	ans := make([]int, 2*len(nums))
	copy(ans[:len(nums)], nums)
    copy(ans[len(nums):], nums)
	return ans
}

func main() {
	nums := []int{2, 3, 4, 5}
	fmt.Println(getConcatenation(nums))
	fmt.Println(getConcatenationByCopying(nums))
}
