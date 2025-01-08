package main

import (
	"fmt"
	"os"
)

func main() {
	myVar := os.Getenv("MY_VAR")
	anotherVar := os.Getenv("ANOTHER_VAR")
	fmt.Println(myVar, anotherVar)
}
