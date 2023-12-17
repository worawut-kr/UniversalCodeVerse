package main

import {
    "fmt"
}
func main() {
   
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }

    // Example 2: for loop as a while loop
    j := 0
    for j < 5 {
        fmt.Println(j)
        j++
    }

    // Example 3: Iterate over elements in a collection (slice)
    fruits := []string{"apple", "banana", "cherry"}
    for index, value := range fruits {
        fmt.Printf("Index: %d, Value: %s\n", index, value)
    }
}
