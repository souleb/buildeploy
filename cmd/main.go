package main

import (
	"github.com/SouleBA/buildeploy/http"
)

func main() {
	server := http.NewServer()
	server.Open()
	defer server.Close()
	for {

	}
}
