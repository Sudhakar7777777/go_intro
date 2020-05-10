package main

import (
	"fmt"
	"math"
)

func sqrt(n float64) (float64, error) {
	if n < 0.0 {
		return 0.0, fmt.Errorf("sqrt of negative value (%f)", n)
	}
	return math.Sqrt(n), nil
}

func main() {

	// go has a concept of panic similar to exception, its discouraged to be used.
	// go suggest to use error as shown in this example

	s1, err := sqrt(2.0)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Println(s1)
	}

	s2, err := sqrt(-2.0)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Println(s2)
	}

}
