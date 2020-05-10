package main

import "fmt"

func main() {
	// slices are built on top of array type in go language.  we seldom use array type directly.

	// create slice of strings
	// same type
	loons := []string{"bugs", "daffy", "taz"}
	fmt.Printf("loons = %v (type %T)\n", loons, loons)

	// length
	fmt.Println(len(loons))

	// 0 indexing
	fmt.Println("-----", loons[1])

	// slices
	fmt.Println("-----", loons[1:])

	// examples to iterate through the slice

	// simple for loop
	for i := 0; i < len(loons); i++ {
		fmt.Println(loons[i])
	}

	// range - single value
	for i := range loons {
		fmt.Println(i)
	}

	// range - double value ; index + value
	for i, name := range loons {
		fmt.Println(i, "::", name)
	}

	// range - double value ; ignore the index using _
	for _, name := range loons {
		fmt.Println(name)
	}

	// add item to existing slice using append
	loons = append(loons, "elmer")
	fmt.Println(loons)
}
