package main

import "fmt"

func main() {
	count := 0

	// for every pair of 4 digit number
	for i := 1000; i <= 9999; i++ {
		for j := i; j <= 9999; j++ { // don't count twice
			num := i * j
			strNum := fmt.Sprintf("%d", num)
			// if strNum is even ended number
			if strNum[0] == strNum[len(strNum)-1] {
				count++
			}
		}
	}

	fmt.Printf("Count = %d", count)
}
