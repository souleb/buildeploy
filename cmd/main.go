package main

import (
	"github.com/souleb/buildeploy/http"
)

func main() {
	server := http.NewServer()
	server.Open()
	defer server.Close()
	for {

	}
}
