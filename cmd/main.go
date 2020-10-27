package main

import (
	"io"
	"os"

	"github.com/souleb/buildeploy/http"
	"github.com/souleb/buildeploy/postgres"
	"github.com/souleb/buildeploy/workflow"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "buildeploydb"
)

const (
	exitFail = 1
)

func main() {

	err := run(os.Stdout)
	if err != nil {
		os.Exit(exitFail)
	}

	//defer client.Close()

	/*ws := postgres.PipelineService{Client: client}
	ws.DestructiveReset()
	ws.Create(&app.Workflow{})

	foundUser, err := ws.GetByID(1)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(foundUser)
	*/

	scheduler := workflow.NewSchedulerService()

	server, err := http.NewServer(scheduler)
	if err != nil {
		panic(err)
	}
	server.Open()
	defer server.Close()
}

func run(stdout io.Writer) error {
	//New db connection
	opt := func(c *postgres.Client) {
		c.Host = host
		c.Port = port
		c.User = user
		c.Password = password
		c.DBname = dbname
	}

	client := postgres.NewClient(opt)
	return client.Open()
}
