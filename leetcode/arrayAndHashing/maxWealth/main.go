package main 

import(
	"fmt"
)

//maxWealth returns the maximum wealth found among all customer
func maximumWealth(accounts [][]int) int {
	maxWealth := 0
	sum := 0
  
	for _, account := range accounts {
	  for i := range account {
		// Add the amount in the current bank account to sum.
		  sum += account[i]
	  }
	  // If the current customer's wealth is greater than maxWealth, update maxWealth.
	  if maxWealth < sum {
		  maxWealth = sum
	  }
	  sum = 0
	}
	return maxWealth
  
}

func main() {
	accounts1 := [][]int{
		{1, 2, 3},
		{3, 2, 1},
		{4, 3, 2},
	}

	fmt.Println("Maximum wealth :", maximumWealth(accounts1)) 

}