package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fileName := "./out"
	f, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	num := 200
	if _, err := fmt.Fprintf(f, "%d\n", num); err != nil {
		log.Fatal(err)
	}

	fmt.Println("file writing has done")
}
