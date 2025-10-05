package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	f, err := os.Create("New.Text")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)

	f2, err := os.Open("File.Txtx")
	if err != nil {

		log.Println("err happend", err)

	}
	defer f2.Close()

}
