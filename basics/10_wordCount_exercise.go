package main

import (
	"fmt"
	"strings"
)

func main() {
	text := `Needles and pins
Needles and pins
Sew me a sail
To catch me the wind`

	words := strings.Fields(text)
	counts := map[string]int{} // create an empty map

	for _, item := range words {
		counts[strings.ToLower(item)]++
	}

	fmt.Println(counts)
}
