package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/omprakash2929/todoaspp/auth"
	"github.com/omprakash2929/todoaspp/user"
)

// ? go mod tidy
// ? go get "packageurl"
func main() {
	auth.LoginWithCredentials("ompraksh", "123456")

	session := auth.Get_session()
	fmt.Println(session)

	user := user.User{
		Email: "ompraksh@gmail.com",
		Name:  "omprakash",
	}

	color.Blue(user.Email)
	color.Yellow(user.Name)
	color.HiRed("Prints text in cyan.")

}
