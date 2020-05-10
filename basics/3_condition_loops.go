package main

import (
	"fmt"
)

func main() {

	fmt.Println("Hello world")

	x := 10
	if x > 5 {
		fmt.Println("X is Bigger than 5")
	}

	if x > 100 {
		fmt.Println("X is too big")
	} else {
		fmt.Println("X is not that big")
	}

	if x > 5 && x < 15 {
		fmt.Println("X is just right")
	}

	a, b := 11.0, 20.0

	if frac := a / b; frac > 0.5 {
		fmt.Printf("a is more than half of b.  Frac=%+v\n", frac)
	}

	switch x {
	case 5:
		fmt.Println("Five")
	case 10:
		fmt.Println("Ten")
	default:
		fmt.Println("I don't know")
	}

	switch {
	case x < 5:
		fmt.Println("X is too small")
	case x > 50:
		fmt.Println("X is too large")
	default:
		fmt.Println("X is right size")
	}

	for i := 1; i <= 5; i++ {
		fmt.Println("I = ", i)
	}

	for i := 1; i <= 5; i++ {
		if i < 2 {
			continue
		} else if i > 3 {
			break
		}
		fmt.Println("CI = ", i)
	}

	j := 0
	for j < 3 {
		fmt.Println("J = ", j)
		j++
	}

	k := 0
	for {
		if k > 2 {
			break
		}
		fmt.Println("K = ", k)
		k++
	}
}
