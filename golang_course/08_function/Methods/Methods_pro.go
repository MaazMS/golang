package main

import "fmt"

type userinfo struct {
	name  string
	email string
	phone int
}

func (user userinfo) show() {

	fmt.Println("User name", user.name)
	fmt.Println("User email", user.email)
	fmt.Println("User phone", user.phone)

}

func main() {

	res := userinfo{
		name:  "Maaz",
		email: "maazshaikh720",
	}

	res.show()
}
