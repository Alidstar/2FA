package main

import (
	"./line"
)

type Server interface {
	Start(port string)
}

const baseURL = "https://abca3074.ngrok.io/"
const auth2Path = "auth2"
const callbackPath = "done"

func main() {
	server, _ := line.NewLineAuthServer(baseURL, auth2Path, callbackPath)
	server.Start("9100")
}
