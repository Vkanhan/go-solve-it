package main

import (
	"fmt"
)

// prime number is only divisible by 1 and the number itself
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	// The loop runs as long as i * i is less than or equal to n.
	for i := 2; i*i <= n; i++ {
		// If n is divisible by i, it means n has a divisor other than 1 and itself.
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	const limit = 200
	fmt.Println("Prime number from 1 to ", limit, ":")

	for i := 2; i <= limit; i++ {
		if isPrime(i) {
			fmt.Println(i, "")
		}
	}
}
