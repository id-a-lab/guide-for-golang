package main

import "fmt"

func main() {

	fruits := [3]string{"apple", "banana", "tomato"}
	for _, fruit := range fruits {
		fmt.Printf("So delicious food, %s\n", fruit)
	}
}
