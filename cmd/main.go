package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/souleb/buildeploy/log"
	"github.com/souleb/buildeploy/pipeline"
	"github.com/souleb/buildeploy/postgres"
	"github.com/souleb/buildeploy/transport"
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
	//reader := bufio.NewReader(os.Stdin)
	//reader.ReadString('\n')
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
	log.Info(fmt.Sprintf("Starting Database connection db=%s, port=%d", dbname, port))
	client := postgres.NewClient(opt)
	if err := client.Open(); err != nil {
		return err
	}
	defer client.Close()
	defer log.Info(fmt.Sprintf("Stopping Database connection db=%s, port=%d", dbname, port))

	ps := &postgres.PipelineService{Client: client}

	log.Info("Starting Listener...")
	server, err := transport.NewServer(ps, log)
	if err != nil {
		return err
	}

	if err = server.Open(); err != nil {
		log.Fatal(err)
	}
	defer server.Close()
	defer log.Info(fmt.Sprintf("Stopping Listener"))

	scheduler := pipeline.NewSchedulerService(log, server)
	log.Info("Starting the Scheduler...")
	err = scheduler.Schedule()
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
