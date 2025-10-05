package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	f1, err := os.Create("names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f1.Close()

	r := strings.NewReader("Maaz shaikh")

	io.Copy(f1, r)

	f2, err := os.Open("names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f2.Close()
	bs, err := ioutil.ReadAll(f2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(bs))
}
