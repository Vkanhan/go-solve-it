package main 

import(
	"fmt"
)

func findDifference(nums1, nums2 []int) [][]int {
	set1 := make(map[int]bool)
	set2 := make(map[int]bool)

	for _, num := range nums1 {
		set1[num] = true 
	}

	for _, num := range nums2 {
		set2[num] = true 
	}

	uniqueNums1 := []int{}
	for num := range set1 {
		if !set2[num] {
			uniqueNums1 = append(uniqueNums1, num)
		}
	}

	uniqueNums2 := []int{}
	for num := range set2 {
		if !set1[num] {
			uniqueNums2 = append(uniqueNums2, num)
		}
	}

	return [][]int{uniqueNums1, uniqueNums2}
}

func main() {
	nums1 := []int{1, 2, 3}
	nums2 := []int{2, 4, 6}

	fmt.Println(findDifference(nums1, nums2))
}