package main 

import(
	"fmt"
)

//hasDuplicate checks if there is a duplicate int in the slice of array
func hasDuplicate(nums []int) bool {
	//Initialize a map to store the unique element encountered
	hashSet := make(map[int]bool)
	
	//Iterate over each element
	for _, num := range nums {
		//Check if the element is already in the map
		if hashSet[num] {
			return true
		}
		//Otherwise add element to the map
		hashSet[num] = true
	}
	//Return false if no duplicates
	return false
}

func main() {
	nums := []int{1, 3, 4, 2, 2}

	fmt.Println(hasDuplicate(nums))
}