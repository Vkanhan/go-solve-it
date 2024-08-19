package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func main() {

	url := "https://codingquest.io/api/puzzledata?puzzle=18"

	//Get the input from the url
	input, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer input.Body.Close()

	// Initialize a map to store the total quantities per category
	categoryTotals := make(map[string]int)

	//Read and parse the data from the input body
	scanner := bufio.NewScanner(input.Body)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		if len(parts) < 3 {
			continue
		}

		quantityStr := parts[1]	//quantity is expected to be the 2 element
		category := parts[2]	//category is expected to be the 3 element.

		//convert the quantity string to integer
		quantity, err := strconv.Atoi(quantityStr)
		if err != nil {
			log.Fatal(err)
		}

		// Add quantity to the specific category total
		categoryTotals[category] += quantity
	}

	//Check if there is any error occured during scanning 
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Calculate modulus 100 for each category total and their product
	product := 1
	for _, total := range categoryTotals {
		mod := total % 100
		product *= mod
	}

	fmt.Printf("Product of moduli: %d\n", product)

	
}