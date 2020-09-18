package main

import (
	"github.com/souleb/buildeploy/http"
	"github.com/souleb/buildeploy/postgres"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "buildeploydb"
)

func main() {

	//New db connection
	opts := func(c *postgres.Client) {
		c.Host = host
		c.Port = port
		c.User = user
		c.Password = password
		c.DBname = dbname
	}

	client := postgres.NewClient(opts)
	err := client.Open()

	if err != nil {
		panic(err)
	}

	//WorkflowService
	//ws := postgres.WorkflowService{Client: client}
	//ws.DestructiveReset()

	server, err := http.NewServer()
	if err != nil {
		panic(err)
	}
	server.Open()
	defer server.Close()
}
