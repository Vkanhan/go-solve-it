package main

import (
	"fmt"
	"strings"
)

func numUniqueEmails(emails []string) int {
	uniqueEmails := make(map[string]struct{}) // Use struct{} as it takes up zero space

	for _, email := range emails {
		// Split the email into local and domain parts
		parts := strings.Split(email, "@")
		if len(parts) != 2 {
			continue // Skip invalid emails
		}
		local, domain := parts[0], parts[1]

		// Remove everything after the first '+' in the local part
		if plusIdx := strings.Index(local, "+"); plusIdx != -1 {
			local = local[:plusIdx]
		}

		// Remove dots from the local part
		local = strings.ReplaceAll(local, ".", "")

		// Combine the processed local and domain parts
		normalizedEmail := local + "@" + domain

		// Store the unique email in the map
		uniqueEmails[normalizedEmail] = struct{}{}
	}

	// The number of unique emails
	return len(uniqueEmails) 
}

func main() {
	emails := []string{
		"test.email+alex@leetcode.com",
		"test.e.mail+bob.cathy@leetcode.com",
		"testemail+david@lee.tcode.com",
	}
	result := numUniqueEmails(emails)
	fmt.Println(result) 
}
