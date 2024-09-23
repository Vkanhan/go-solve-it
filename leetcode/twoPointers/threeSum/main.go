package main

import (
	"fmt"
	"sort"
)

// threeSum finds all unique triplets in the array that sum to zero.
func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	var result [][]int

	for i := 0; i < len(nums); i++ {
		//If the current number is greater than zero, stop the loop as further numbers can't sum to zero.
		if nums[i] > 0 {
			break
		}

		//Skip this number if it's the same as the previous one to avoid duplicate triplets.
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		//Use the two-pointer technique to find two numbers that sum to the negative of the current number.
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if sum > 0 {
				right--
			} else if sum < 0 { 
				left++
			} else {
				// If the sum is exactly zero, we've found a valid triplet.
				// Add the triplet to the result slice.
				result = append(result, []int{nums[i], nums[left], nums[right]})

				left++
				right--

				//Skip duplicate numbers on the left side to avoid adding the same triplet multiple times.
				for left < right && nums[left] == nums[left-1] {
					left++
				}
			}
		}
	}
	return result
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}

	fmt.Println(threeSum(nums)) 
}
