package main 

import(
	"fmt"
)

//hasDuplicate checks if there is a duplicate int in the slice of array
func hasDuplicate(nums []int) bool {
	hashSet := make(map[int]bool)
	
	for _, num := range nums {
		if hashSet[num] {
			return true
		}
		hashSet[num] = true
	}
	return false
}

func main() {
	nums := []int{1, 3, 4, 2, 2}

	fmt.Println(hasDuplicate(nums))
}