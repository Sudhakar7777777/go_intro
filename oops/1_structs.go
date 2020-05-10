package main

import (
	"fmt"
	"os"
)

type Trade struct {
	Symbol string  // Stock symbol
	Volume int     // Number of shares
	Price  float64 // Trade Price
	Buy    bool    // true if buy trade, false if sell trade
}

// defining methods for struct
// value returns the trade value
func (t *Trade) Value() float64 {
	value := float64(t.Volume) * t.Price
	if t.Buy {
		value = -value
	}
	return value
}

// defining a constructor
func NewTrade(symbol string, volume int, price float64, buy bool) (*Trade, error) {
	if symbol == "" {
		return nil, fmt.Errorf("Symbol can't be empty")
	}

	if volume <= 0 {
		return nil, fmt.Errorf("Volume must be greater than 0.  (was %d)", volume)
	}

	if price <= 0 {
		return nil, fmt.Errorf("Volume must be greater than 0.  (was %f)", price)
	}

	t := &Trade{
		Symbol: symbol,
		Volume: volume,
		Price:  price,
		Buy:    buy,
	}

	return t, nil
}

func main() {
	// define custom data type in go using struct

	t1 := Trade{"MSFT", 90, 160.50, true} // pass values in the same order its defined
	fmt.Println(t1)

	fmt.Printf("%v\n", t1)

	fmt.Printf("%+v\n", t1)

	fmt.Println(t1.Symbol)

	t2 := Trade{ // if using key value pairs, no order is required
		Symbol: "APPL",
		Buy:    false,
		Price:  830.23,
		Volume: 1000,
	}
	fmt.Printf("%+v\n", t2)

	t3 := Trade{}
	fmt.Printf("%+v\n", t3)

	// in go everything that starts with an upper case letter is accessible from other packages
	// otherwise its only accessible within current package

	// calling struct methods
	fmt.Println("Value of Trade is : ", t1.Value())

	// create new object using the constructor method
	t4, err := NewTrade("SBK", 100, 7777.77, true)
	if err != nil {
		fmt.Printf("Error: Creating Trade instance.  %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Trade object using constructor. T4 = %+v\n", t4)

}
