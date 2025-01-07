package main

import (
	"fmt"
)

var fruits []string = []string{"apple", "banana", "cherry", "tomato", "avacodo"}

func main() {
	// for _, fruits := range fruits {
	// 	fmt.Println(fruits)
	// }
	for _, value := range fruits[:2] {
		fmt.Println(value)
	}
}
