package main

import (
    "fmt"
    "strconv"
    "strings"
)

// Encode converts a list of strings into a single encoded string.
func Encode(strs []string) string {
    var encoded strings.Builder
    for _, s := range strs {
        encoded.WriteString(strconv.Itoa(len(s)))
        encoded.WriteString("#")
        encoded.WriteString(s)
    }
    return encoded.String()
}

// Decode converts an encoded string back into a list of strings.
func Decode(s string) []string {
    var result []string
    i := 0
    for i < len(s) {
        // Find the position of the '#' character
        j := i
        for s[j] != '#' {
            j++
        }
        
        // Extract the length of the string
        length, _ := strconv.Atoi(s[i:j])
        
        // Move past the '#' character
        j++
        
        // Extract the actual string based on the length
        result = append(result, s[j:j+length])
        
        // Move the index past the current string
        i = j + length
    }
    return result
}

func main() {
    strs := []string{"hello", "world"}
    encoded := Encode(strs)
    fmt.Println("Encoded:", encoded)
    
    decoded := Decode(encoded)
    fmt.Println("Decoded:", decoded)
}
