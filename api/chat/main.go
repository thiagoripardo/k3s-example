package main

import (
	"github.com/thiagoripardo/k3s-example/api/chat/server"
)

func main() {
	err := server.New().Run()
	if err != nil {
		panic(err)
	}
}
