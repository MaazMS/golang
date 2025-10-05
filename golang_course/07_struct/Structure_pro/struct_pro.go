package main

import "fmt"

func main() {

	type userinfo struct {
		name  string
		email string
		phone int64
	}

	user1 := userinfo{"Maaz", "maazshaikh7020@gmail.com", 123}
	fmt.Println(user1)

	user2 := userinfo{name: "Shaikh", email: "maazshaikh170@gmail.com", phone: 123}
	fmt.Println(user2)
}
