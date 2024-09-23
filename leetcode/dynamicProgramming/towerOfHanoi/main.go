package main

import "fmt"

func towerOfHanoi(n int, source string, destination string, helper string) {
    if n == 1 {
        fmt.Printf("Move disk 1 from %s to %s\n", source, destination)
        return
    }
    
    // Move n-1 disks from source to helper using destination as a temporary peg
    towerOfHanoi(n-1, source, helper, destination)
    
    // Move the nth disk from source to destination
    fmt.Printf("Move disk %d from %s to %s\n", n, source, destination)
    
    // Move n-1 disks from helper to destination using source as a temporary peg
    towerOfHanoi(n-1, helper, destination, source)
}

func main() {
    n := 3 
    towerOfHanoi(n, "A", "C", "B") 
}
