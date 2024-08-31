package main

import "fmt"

// productExceptSelf returns an array where each element is the product of all elements 
//in the input array except the element at that index.
func productExceptSelf(nums []int) []int {
    n := len(nums)
    result := make([]int, n)
    
	// prefix store the cumulative product to the left of the current index
    prefixProduct := 1

    for i := 0; i < n; i++ {
        result[i] = prefixProduct	// Set the current index of result to the current prefix value
        prefixProduct *= nums[i]	// Update the prefix to include the current element
    }
    
	//postfix store the cumulative product to the right of the current index
    postfixProduct := 1
    for i := n - 1; i >= 0; i-- {
        result[i] *= postfixProduct
        postfixProduct *= nums[i]
    }
    
    return result
}

//Brute force approach
func ProductExceptSelf(nums []int) []int {
    n := len(nums)
    result := make([]int, n)

    for i := 0; i < n; i++ {
        product := 1
        for j := 0; j < n; j++ {
            if i != j {
                product *= nums[j]
            }
            result[i] = product
        }
    }
    return result
}

func productExceptSelfByDivision(nums []int) []int {
    n := len(nums)
    result := make([]int, n)

    // Step 1: Calculate the product of all elements
    totalProduct := 1
    for i := 0; i < n; i++ {
        totalProduct *= nums[i]
    }

    // Step 2: Fill the result array by dividing totalProduct by each element
    for i := 0; i < n; i++ {
        result[i] = totalProduct / nums[i]
    }
    return result
}

func main() {
    nums := []int{1, 2, 3, 4}
    fmt.Println(productExceptSelf(nums))
    fmt.Println(ProductExceptSelf(nums))
    fmt.Println(productExceptSelfByDivision(nums))
}
