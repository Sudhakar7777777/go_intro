// function that gets the URL and returns the value of Content-Type response HTTP Header
// function to return error if the Get request fails
package main

import (
	"fmt"
	"net/http"
)

func contentType(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close() // make sure we close the body

	ctype := resp.Header.Get("Content-Type")
	if ctype == "" { // return error if content type not found
		return "", fmt.Errorf("Content-Type header not found")
	}
	return ctype, nil
}

func main() {
	ctype, err := contentType("https://golang.org")
	if err != nil {
		fmt.Printf("Error : %s\n", err)
	} else {
		fmt.Println(ctype)
	}

	ctype, err = contentType("https://sudhakar.org")
	if err != nil {
		fmt.Printf("Error : %s\n", err)
	} else {
		fmt.Println(ctype)
	}
}
