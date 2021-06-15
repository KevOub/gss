package main

import (
	"fmt"

	encoder "github.com/KevOub/gss/encoder"
)

func main() {

	test := encoder.NewEncoder("assets/cat.png")
	fmt.Printf("PATH %s", test)

	// test.Path
	fmt.Println("hello world")
}
