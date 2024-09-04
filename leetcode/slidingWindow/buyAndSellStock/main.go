package main 

import(
	"fmt"
)

func bestTimeToBuyAndSell(prices []int) int {
    if len(prices) == 0 {
        return 0 
    }
    
    maxProfit := 0 
    buyingPrice := prices[0] 

    for _, price := range prices {
        if price < buyingPrice {
            buyingPrice = price 
        } else { 
            profit := price - buyingPrice
            
            if profit > maxProfit {
                maxProfit = profit
            }
        }
    }
    return maxProfit
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(bestTimeToBuyAndSell(prices)) 
}