package auth

import (
	"github.com/fatih/color"
)

func LoginWithCredentials(username string, password string) {
	color.Green("Login user using", username, password)
}
