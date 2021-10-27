package main

import (
	"fmt"
	"os"
)

func main() {

	fruits := make([]string, 3)
	fruits[0] = "apple"
	fruits[1] = "banana"
	fruits[2] = "orange"

	for i := 1; i < len(os.Args); i++ {

		fruits = append(fruits[:len(fruits)], os.Args[i])
	}
	fmt.Println(fruits)
}
