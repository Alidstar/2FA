package main

import (
	"fmt"

	"./line"
	"./um"
)

type Server interface {
	Start(port string)
}

const baseURL = "https://abca3074.ngrok.io/"
const auth2Path = "auth2"
const callbackPath = "done"

func main() {
	um := um.NewUserManager()

	username := "user1"
	um.Register(username)
	token := um.Login(username)
	fmt.Println("Token", token)
	result, username2 := um.Verify(token)
	fmt.Println("Verify:", username2, result)

	if false {
		server, _ := line.NewLineAuthServer(baseURL, auth2Path, callbackPath)
		server.Start("9100")
	}
}
