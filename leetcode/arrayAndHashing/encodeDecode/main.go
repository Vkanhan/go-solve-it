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
        j := i
        for s[j] != '#' {
            j++
        }
        
        length, _ := strconv.Atoi(s[i:j])
        
        j++
        
        result = append(result, s[j:j+length])
        
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
