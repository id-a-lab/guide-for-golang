package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	fileName := "./in"
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var value int
	if _, err := fmt.Fscanf(f, "%d\n", &value); err != nil {
		log.Fatal(err)
	}

	outFileName := fmt.Sprintf("./out_%d", value)
	if f, err = os.Create(outFileName); err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if _, err = fmt.Fprintf(f, "there is no contents"); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Process has done!")

}
