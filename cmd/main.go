package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/souleb/buildeploy/log"

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

	debug := flag.Bool("debug", false, "sets log level to debug")

	flag.Parse()

	level := "info"

	if *debug {
		level = "debug"
	}

	log := log.New(os.Stdout, level)

	err := run(os.Stdout, log)
	if err != nil {
		log.Fatal(err)
	}

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
}

func run(stdout io.Writer, log *log.Logger) error {
	//New db connection
	opt := func(c *postgres.Client) {
		c.Host = host
		c.Port = port
		c.User = user
		c.Password = password
		c.DBname = dbname
	}
	log.Info(fmt.Sprintf("Starting Database db=%s, port=%d", dbname, port))
	client := postgres.NewClient(opt)
	if err := client.Open(); err != nil {
		return err
	}
	defer client.Close()
	defer log.Info(fmt.Sprintf("Stopping Database db=%s, port=%d", dbname, port))

	ps := &postgres.PipelineService{Client: client}
	scheduler := workflow.NewSchedulerService()

	log.Info(fmt.Sprintf("Starting Listener"))
	server, err := http.New(scheduler, ps)
	if err != nil {
		return err
	}
	if err = server.Open(); err != nil {
		return err
	}

	defer server.Close()
	defer log.Info(fmt.Sprintf("Stopping Listener"))

	return nil
}
