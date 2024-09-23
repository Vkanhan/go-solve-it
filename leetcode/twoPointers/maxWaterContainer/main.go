package main 

import(
	"fmt"
)

//maxArea stores the maximum amount of water a container can store
func maxArea(height []int) int {
	//Initialize two pointers
	left, right := 0, len(height) - 1

	maxArea := 0

	for left < right {
		//Calculate the width and height of the container formed by the two pointer
		width := right - left

		h := min(height[left], height[right])

		//Calculate the area and update maxArea if the current area is larger
		currentArea := width * h
		if currentArea > maxArea {
			maxArea = currentArea
		}

		//Move the pointer pointing to the shorter height inward
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}

//Helper function to find the minimum
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	height1 := []int{1, 7, 2, 5, 4, 7, 3, 6}
	fmt.Println("Maximum water container for height1:", maxArea(height1)) 

	height2 := []int{2, 2, 2}
	fmt.Println("Maximum water container for height2:", maxArea(height2)) 
}