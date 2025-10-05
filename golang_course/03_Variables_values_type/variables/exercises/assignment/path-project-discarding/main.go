package main

import (
	"fmt"
	"path"
)

func main()  {

	dir , file := path.Split("/home/maaz/composer.json")
	fmt.Println("dir:", dir)
	fmt.Println("file:", file)
}
