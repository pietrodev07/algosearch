package main

import (
	"algosearch/searching"
	"fmt"
)

func main() {
	arr := []int{10, 20, 30, 40, 50}
	target := 30

	// Linear Search
	fmt.Printf("Array: %v\n", arr)
	fmt.Printf("Linear Search for %d: index %d\n", target, searching.LinearSearch(arr, target))

	// Binary Search
	fmt.Printf("Binary Search for %d: index %d\n", target, searching.BinarySearch(arr, target))
}
