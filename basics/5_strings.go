package main

import (
	"fmt"
)

func main() {
	book := "The color of magic!!"

	fmt.Println(book)
	fmt.Println(len(book))
	fmt.Printf("book[0]=%v (type %T)\n", book[0], book[0])
	//uint8 is byte
	//Strings in go are immutable; can't do.... book[0] = 116

	//Slice full form
	fmt.Println(book[4:10])
	//Slice no end
	fmt.Println(book[4:])
	//Slice no start
	fmt.Println(book[:4])

	//concatenate
	fmt.Println("!!" + book)

	//unicode by default
	fmt.Println("One ½ years old.")

	//multiline
	poem := `
	Twinkle Twinkle Little Star.
	How I wonder what you are?
	Up above the world so high.
	Like a diamon in the sky.
	`
	fmt.Println(poem)

	num := 45
	strNum := fmt.Sprintf("%d", num)
	fmt.Printf("strNum = %q (type %T)\n", strNum, strNum)

}
