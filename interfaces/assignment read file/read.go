package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	filename := os.Args[1]

	//to open the file
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}


	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("read %d bytes: %q\n", count, data[:count])

}