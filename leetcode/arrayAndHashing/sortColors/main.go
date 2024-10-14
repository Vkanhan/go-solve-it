package main

import "fmt"

// sortColors sorts an array in place, where the array contains only 0s, 1s, and 2s, representing different colors
func sortColor(colorArray []int) {
	countOfZeros, countOfOnes := 0, 0
	totalLength := len(colorArray)

	// Count the number of 0s and 1s
	for _, color := range colorArray {
		if color == 0 {
			countOfZeros++
		} else if color == 1 {
			countOfOnes++
		}
	}

	//Fill the array with 0
	for i := 0; i < countOfZeros; i++ {
		colorArray[i] = 0
	}

	//Fill the array with 1
	for i := countOfZeros; i < countOfZeros+countOfOnes; i++ {
		colorArray[i] = 1
	}

	//Fill the array with 2
	for i := countOfOnes + countOfOnes; i < totalLength; i++ {
		colorArray[i] = 2
	}
}

func main() {
	colors := []int{2, 1, 0, 1, 0, 2}
	fmt.Println("Original colors: ", colors)

	sortColor(colors)

	fmt.Println("Sorted colors: ", colors)
}
