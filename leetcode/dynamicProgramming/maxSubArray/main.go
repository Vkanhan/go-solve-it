package main

import "fmt"

//find the largest non empty subarray
func largestSubarray(nums []int) int {
	maxSum := nums[0]

	for i := 0; i < len(nums); i++ {
		currSum := 0
		for j := i; j < len(nums); j++ {
			currSum += nums[j]
			if maxSum < currSum {
				maxSum = currSum
			}
		}
	}
	return maxSum
}

func kadanes(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxSum := nums[0]
	currSum := 0

	for _, n := range nums {
		if currSum < 0 {
			currSum = 0
		}
		currSum += n

		if currSum > maxSum {
			maxSum = currSum
		}
	}
	return maxSum
}

func main() {
	nums := []int{2, 3, 2, 4, 7, -5}
	fmt.Println(largestSubarray(nums))
	fmt.Println(kadanes(nums))
}

