package main 

import(
	"fmt"
)

//intersection returns an array that contains all the unique elements that are present in both arrays
func intersection(nums1 []int, nums2 []int) []int {
    result := []int{}

    hashSet := make(map[int]bool)
    resultSet := make(map[int]bool)  // To track elements already added to the result
    
    // Populate hashSet with elements from nums1
    for _, val := range nums1 {
        hashSet[val] = true
    }

    // Check for intersection and add to result
    for _, val := range nums2 {
        if hashSet[val] && !resultSet[val] {  // Ensure the element hasn't already been added to result
            result = append(result, val)
            resultSet[val] = true
        }
    }
    
    return result
}


func main() {
	nums1 := []int{1, 2, 3, 4, 5}

	nums2 := []int{4, 5, 6, 7, 8}

	fmt.Println(intersection(nums1, nums2))
}