package main

import (
	"fmt"
	"os"

	"github.com/souleb/buildeploy/app"
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
	opt := func(c *postgres.Client) {
		c.Host = host
		c.Port = port
		c.User = user
		c.Password = password
		c.DBname = dbname
	}

	client := postgres.NewClient(opt)
	err := client.Open()

	if err != nil {
		panic(err)
	}

	//WorkflowService
	ws := postgres.WorkflowService{Client: client}
	ws.DestructiveReset()
	ws.Create(&app.Workflow{})

	foundUser, err := ws.GetByID(1)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(foundUser)

	server, err := http.NewServer()
	if err != nil {
		panic(err)
	}
	server.Open()
	defer server.Close()
}
