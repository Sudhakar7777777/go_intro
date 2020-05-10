//program to
// print 1 to 20
// Print fizz if number divisible by 3
// Print buzz if number divisible by 5
// Print fizz buzz if number divisible by 3 & 5
package main

import (
	"fmt"
)

func main() {

	for i := 1; i <= 20; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("fizz buzz")
		case i%3 == 0:
			fmt.Println("fizz")
		case i%5 == 0:
			fmt.Println("buzz")
		default:
			fmt.Println(i)
		}
	}
}
