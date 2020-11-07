package postgres

import (
	"context"
	"database/sql/driver"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/souleb/buildeploy/app"
)

// Ensure PipelineService implements app.PipelineService.
var _ app.PipelineService = (*PipelineService)(nil)

// PipelineService represents a PostgreSQL implementation of app.PipelineService.
type PipelineService struct {
	Client *Client
}

type autoIncr struct {
	id        int64
	createdAt time.Time
	updatedAt time.Time
	deletedAt time.Time
}

// workflow is used for the db representation of a workflow
type workflow struct {
	name       string
	pipelineID int64
	autoIncr
}

// Commands are actions to be performed
type commands []string

func (commands commands) Value() (driver.Value, error) {
	var quotedStrings []string
	for _, str := range commands {
		quotedStrings = append(quotedStrings, strconv.Quote(str))
	}

	Values := fmt.Sprintf("{ %s }", strings.Join(quotedStrings, ","))

	return Values, nil
}

func (commands *commands) Scan(src interface{}) error {
	val, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("unable to scan")
	}
	Values := strings.TrimPrefix(string(val), "{")
	Values = strings.TrimSuffix(Values, "}")

	*commands = strings.Split(Values, ",")

	return nil
}

// job is used for the db representation of a job
type job struct {
	name       string
	workflowID int64
	edges      string
	steps      commands
	env        string
	branches   string
	status     app.Status
	autoIncr
}

type docker struct {
	image string
	tags  string
	jobID int64
	autoIncr
}

type machine struct {
	os     string
	cpus   string
	memory string
	jobID  int64
	autoIncr
}

// ReadWorkflow will get a workflow by ID.
//func (c *Client) ReadWorkflow(ctx context.Context, id uint, workflow *app.Workflow) error {
//	err := c.DB.GetContext(ctx, &workflow, "SELECT * FROM workflow WHERE id == $1", id)
//	if err != nil {
//		return errors.Wrap(err, "sql: ID provided was invalid")
//	}
//
//	return nil
//}

// GetJobByID will look up a pipeline with the provided ID.
// If the Job is found, we will return a nil error
// If there is an error, we will return an error with
// more information about what went wrong.
func (p *PipelineService) GetJobByID(ctx context.Context, id int64) (*app.Job, error) {
	params := QueryParams{
		Query: "SELECT * FROM job WHERE id == $1",
		ID:    id,
		Value: app.Job{},
	}
	err := p.Client.ReadByID(ctx, &params)
	if err != nil {
		return nil, errors.Wrap(err, "failed retrieving the pipeline")
	}

	res := params.Value.(*app.Job)
	return res, nil
}

// GetPipelineByID will look up a pipeline with the provided ID.
// If the pipeline is found, we will return a nil error
// If there is an error, we will return an error with
// more information about what went wrong.
func (p *PipelineService) GetPipelineByID(ctx context.Context, id int64) (*app.Pipeline, error) {
	params := QueryParams{
		Query: "SELECT * FROM pipeline WHERE id == $1",
		ID:    id,
		Value: app.Pipeline{},
	}
	err := p.Client.ReadByID(ctx, &params)
	if err != nil {
		return nil, errors.Wrap(err, "failed retrieving the pipeline")
	}

	res := params.Value.(*app.Pipeline)
	return res, nil
}

// CreatePipeline will create the provided pipeline and backfill data
// like the ID, CreatedAt, and UpdatedAt fields.
func (p *PipelineService) CreatePipeline(ctx context.Context, pipeline *app.Pipeline) (int64, error) {
	var args []interface{}
	args = append(args, pipeline.Name)
	args = append(args, pipeline.Status)
	args = append(args, pipeline.CreatedAt)
	//args = append(args, pipeline.DeletedAt)
	params := ExecParams{
		InsertCmd: "INSERT INTO pipeline(name, status, created_at) VALUES($1, $2, $3) RETURNING id",
		Values:    args,
	}

	return p.Client.Create(ctx, &params)
}

// createWorkflow will create the provided workflow and backfill data
// like the ID, CreatedAt, and UpdatedAt fields.
func (p *PipelineService) createWorkflow(ctx context.Context, workflow *workflow) (int64, error) {
	var args []interface{}
	args = append(args, workflow.name)
	args = append(args, workflow.pipelineID)
	params := ExecParams{
		InsertCmd: "INSERT INTO workflow(name, pipeline_id) VALUES($1, $2) RETURNING id",
		Values:    args,
	}

	return p.Client.Create(ctx, &params)
}

// createJob will create the provided job and backfill data
// like the ID, CreatedAt, and UpdatedAt fields.
func (p *PipelineService) createJob(ctx context.Context, job *job) (int64, error) {
	var args []interface{}
	args = append(args, job.name)
	args = append(args, job.workflowID)
	args = append(args, job.edges)
	args = append(args, job.steps)
	args = append(args, job.env)
	args = append(args, job.branches)
	args = append(args, job.status)
	params := ExecParams{
		InsertCmd: "INSERT INTO job(name, workflow_id, edges, steps, env, branches, status" +
			") VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING id",
		Values: args,
	}

	return p.Client.Create(ctx, &params)
}

// createDocker will create the provided docker and backfill data
// like the ID, CreatedAt, and UpdatedAt fields.
func (p *PipelineService) createDocker(ctx context.Context, docker *docker) (int64, error) {
	var args []interface{}
	args = append(args, docker.jobID)
	args = append(args, docker.image)
	args = append(args, docker.tags)
	params := ExecParams{
		InsertCmd: "INSERT INTO job_docker(job_id, image, tags) VALUES($1, $2, $3) RETURNING id",
		Values:    args,
	}

	return p.Client.Create(ctx, &params)
}

// createMachine will create the provided machine and backfill data
// like the ID, CreatedAt, and UpdatedAt fields.
func (p *PipelineService) createMachine(ctx context.Context, machine *machine) (int64, error) {
	var args []interface{}
	args = append(args, machine.os)
	args = append(args, machine.jobID)
	args = append(args, machine.cpus)
	args = append(args, machine.memory)
	params := ExecParams{
		InsertCmd: "INSERT INTO job_machine(os, job_id, cpus, memory) VALUES($1, $2, $3, $4) RETURNING id",
		Values:    args,
	}

	return p.Client.Create(ctx, &params)
}
