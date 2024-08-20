package main 

import(
	"fmt"
)
//Floys cycle finding algo also known as tortoise-hare algo 
//to find the duplicate number in an array when the elements are in the range [1, n] and exactly one duplicate number
func tortoiseHare(nums []int) int {
	//Initialize both pointer to the first element
	tortoise, hare := nums[0], nums[0]

	//find the meeting point of 2 runners
	for {
		tortoise, hare = nums[tortoise], nums[nums[hare]] //move tortoise one step and hare 2 step 
		if tortoise == hare {break}
	}


	//find the entrance to the cycle that is duplicate number
	start, meetingPoint := nums[0], tortoise
	for start != meetingPoint {
		start, meetingPoint = nums[start], nums[meetingPoint] //move start and meetingPoint one step
	}

	//return the duplicate number
	return start
}

func main() {
	nums := []int{1, 3, 4, 2, 2}

	fmt.Println(tortoiseHare(nums))
}