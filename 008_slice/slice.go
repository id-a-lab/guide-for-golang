package main

import "fmt"

func main() {

	fruits := make([]string, 3)
	fruits[0] = "apple"
	fruits[1] = "banana"
	fruits[2] = "tomato"

	for _, fruit := range fruits {
		fmt.Printf("%s\n", fruit)
	}
	fmt.Println("Append grape")
	fmt.Println("================================")

	fruits = append(fruits, "grape")

	for _, fruit := range fruits {
		fmt.Printf("%s\n", fruit)
	}

}
