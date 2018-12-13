package main

import (
	"fmt"

	"./line"
	tm "./token_manager"
)

type Server interface {
	Start(port string)
}

const baseURL = "https://abca3074.ngrok.io/"
const auth2Path = "auth2"
const callbackPath = "done"

func main() {
	manager := tm.NewTokenManager()

	username := "user1"
	token := manager.GenerateToken(username)
	fmt.Println("Token", token)
	result, username2 := manager.Verify(token)
	fmt.Println("Verify:", username2, result)

	if false {
		server, _ := line.NewLineAuthServer(baseURL, auth2Path, callbackPath)
		server.Start("9100")
	}
}
