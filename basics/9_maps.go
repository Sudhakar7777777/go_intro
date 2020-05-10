package main

import "fmt"

func main() {
	// Map is Key Value pairs
	// In go the keys should be of same type and values should be of same type

	stocks := map[string]float64{
		"AMZN": 1699.8,
		"GOOG": 1129.19,
		"MSFT": 98.61, // must have trailing comma in multi line
	}

	fmt.Println("Map = ", stocks)

	// number of items
	fmt.Println(len(stocks))

	// print an item value
	fmt.Println(stocks["GOOG"])

	// get 0 value if item not found
	fmt.Println(stocks["Sudhakar"])

	// use 2 values to find if the actual value is returned; ok has true or false based on item found
	value, ok := stocks["SBK"]
	if ok {
		fmt.Println("Found value :", value)
	} else {
		fmt.Println("Value not found for key", "SBK")
	}

	// set or add a new item
	stocks["SBK"] = 999.99
	fmt.Println(stocks)

	// delete an existing item
	delete(stocks, "AMZN")
	fmt.Println(stocks)

	// iterate for all keys
	for key := range stocks {
		fmt.Println(key)
	}

	// iterate for both key + value
	for key, value := range stocks {
		fmt.Printf("key = %v, value = %v\n", key, value)
	}
}
