package main 

import(
	"fmt"
)

// Function to check if there exists a pair with the given sum in a sorted array
func hasPairWithSum(arr []int, target int) bool {
	// Initialize two pointers
	left, right := 0, len(arr) - 1

	for left < right {
		// Calculate sum of the elements at pointers
		sum := arr[left] + arr[right]
		if sum == target {
			return true
		} else if sum < target {
			left++	// Move left pointer to the right to increase sum
		} else {
			right--	// Move right pointer to the left to decrease sum
		}
	}
	return false
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
    target := 10
    fmt.Println(hasPairWithSum(arr, target))
}