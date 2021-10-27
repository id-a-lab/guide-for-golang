package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fileName := "./abc"
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer f.Close()

	var num int
	if _, err := fmt.Fscanf(f, "%d\n", &num); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Read value is %d", num)
}
