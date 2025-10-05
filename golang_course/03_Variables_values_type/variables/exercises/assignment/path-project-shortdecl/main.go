package main

import (
	"fmt"
	"path"
)

func main()  {

	_, file := path.Split("/home/maaz/composer.json")
	fmt.Println("file:", file)
}
