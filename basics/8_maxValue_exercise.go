// calculate max value in a slice of integers
package main

import "fmt"

func main() {
	nums := []int{16, 8, 42, 4, 23, 15}

	max := nums[0]                   // initialize with first value of slice
	for _, value := range nums[1:] { // skip first value
		if max < value {
			max = value
		}
	}

	fmt.Println("Max value is ", max)
}
