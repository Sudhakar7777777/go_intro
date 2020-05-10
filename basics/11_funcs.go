package main

import "fmt"

// returns integer sum of 2 input values
func add(a int, b int) int {
	return a + b
}

// returns integer quotient and reminder of 2 input values
func divmod(a int, b int) (int, int) {
	return a / b, a % b
}

func doubleAt(values []int, at int) {
	values[at] *= 2
}

func double(n int) {
	n *= 2
}

func doublePtr(n *int) {
	*n *= 2
}

func main() {
	val := add(4, 5)
	fmt.Println("Sum value is ", val)

	div, mod := divmod(7, 2)
	fmt.Printf("Quotient/div = %d, Reminder/mod = %d\n", div, mod)

	// if function paramater type is slice or map by default, the value will be passed by reference
	vals := []int{1, 2, 3, 4}
	doubleAt(vals, 3)
	fmt.Println(vals)

	// if function paramater type is standard type like int or float, etc. by default, the value will be passed by value
	num := 10
	double(num)
	fmt.Println(num)

	// if you want to pass integer to function as pass by reference use the pointers method
	doublePtr(&num)
	fmt.Println(num)
}
