package main 

import(
	"fmt"
)

//removeDuplicates removes the duplicate elements in the slice
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	//Pointer for the index of unique elements
	uniqueIndex := 0

	for i := 0; i < len(nums); i++ {
		//if the current element is different from the last unique element
		if nums[i] != nums[uniqueIndex] {
			uniqueIndex++ //move the uniqueIndex 
			nums[uniqueIndex] = nums[i] //update the position with the new unique element
		}
	}
	return uniqueIndex + 1 //return the length of the unique element
}

func main() {
    nums := []int{1, 1, 2, 2, 3, 4, 4}
    length := removeDuplicates(nums)
    fmt.Println(nums[:length]) 
}