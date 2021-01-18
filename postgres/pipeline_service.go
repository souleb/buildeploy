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
	Log    app.LoggerService
}

type autoIncr struct {
	id        int64
	createdAt time.Time
	updatedAt time.Time
	deletedAt time.Time
}

type pipeline struct {
	name   string
	status app.Status
	autoIncr
}

// workflow is used for the db representation of a workflow
type workflow struct {
	name       string
	pipelineID int64
	autoIncr
}

type array []string

func (array array) Value() (driver.Value, error) {
	var quotedStrings []string
	for _, str := range array {
		quotedStrings = append(quotedStrings, strconv.Quote(str))
	}

	values := fmt.Sprintf("{ %s }", strings.Join(quotedStrings, ","))

	return values, nil
}

func (array *array) Scan(src interface{}) error {
	val, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("unable to scan")
	}
	values := strings.TrimPrefix(string(val), "{")
	values = strings.TrimSuffix(values, "}")

	*array = strings.Split(values, ",")

	return nil
}

// job is used for the db representation of a job
type job struct {
	name       string
	workflowID int64
	needs      array
	steps      array
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

// GetJobByID will look up a pipeline with the provided ID.
// If the Job is found, we will return a nil error
// If there is an error, we will return an error with
// more information about what went wrong.
func (p *PipelineService) GetJobByID(ctx context.Context, id int64) (*app.Job, error) {
	params := queryParams{
		query: "SELECT * FROM job WHERE id == $1",
		id:    id,
		value: app.Job{},
	}
	err := p.Client.readByID(ctx, &params)
	if err != nil {
		return nil, errors.Wrap(err, "failed retrieving the pipeline")
	}

	res := params.value.(*app.Job)
	return res, nil
}

// GetPipelineByID will look up a pipeline with the provided ID.
// If the pipeline is found, we will return a nil error
// If there is an error, we will return an error with
// more information about what went wrong.
func (p *PipelineService) GetPipelineByID(ctx context.Context, id int64) (*app.Pipeline, error) {
	params := queryParams{
		query: "SELECT * FROM pipeline WHERE id == $1",
		id:    id,
		value: app.Pipeline{},
	}
	err := p.Client.readByID(ctx, &params)
	if err != nil {
		return nil, errors.Wrap(err, "failed retrieving the pipeline")
	}

	res := params.value.(*app.Pipeline)
	return res, nil
}

// CreatePipeline will create the provided pipeline and backfill data
// like the ID, CreatedAt, and UpdatedAt fields.
func (p *PipelineService) CreatePipeline(ctx context.Context, appPipeline *app.Pipeline) error {

	pID, err := p.createPipeline(ctx, &pipeline{name: appPipeline.Name, status: appPipeline.Status})
	if err != nil {
		return errors.Wrap(err, "Could not create the pipeline.")
	}
	appPipeline.ID = pID

	for k, w := range appPipeline.Workflows {
		wID, err := p.createWorkflow(ctx, &workflow{name: w.Name, pipelineID: pID})
		if err != nil {
			return errors.Wrap(err, "Could not create the pipeline's workflow")
		}
		appPipeline.Workflows[k].ID = wID

		for k, j := range w.Jobs {
			jID, err := p.createJob(ctx, &job{name: j.Name, workflowID: wID, needs: j.Needs, steps: array(j.Steps),
				env: j.Env, branches: j.Branches, status: j.Status})
			if err != nil {
				return errors.Wrap(err, "Could not create the workflow's job")
			}
			w.Jobs[k].ID = jID
			switch j.Runner.(type) {
			case *app.Docker:
				dID, err := p.createDocker(ctx, &docker{image: j.Runner.(*app.Docker).Image, tags: j.Runner.(*app.Docker).Tags, jobID: jID})
				if err != nil {
					return errors.Wrap(err, "Could not create the job's docker runner")
				}
				j.Runner.(*app.Docker).ID = dID
			case *app.Machine:
				mID, err := p.createMachine(ctx, &machine{os: j.Runner.(*app.Machine).OS, cpus: j.Runner.(*app.Machine).Cpus,
					memory: j.Runner.(*app.Machine).Memory, jobID: jID})
				if err != nil {
					return errors.Wrap(err, "Could not create the job's machine runner")
				}
				j.Runner.(*app.Machine).ID = mID
			}
		}
	}

	return nil
}

func (p *PipelineService) createPipeline(ctx context.Context, pipeline *pipeline) (int64, error) {
	var args []interface{}
	args = append(args, pipeline.name)
	args = append(args, pipeline.status)
	params := execParams{
		insertCmd: "INSERT INTO pipeline(name, status) VALUES($1, $2) RETURNING id",
		values:    args,
	}

	return p.Client.create(ctx, &params)
}

// createWorkflow will create the provided workflow and backfill data
// like the ID, CreatedAt, and UpdatedAt fields.
func (p *PipelineService) createWorkflow(ctx context.Context, workflow *workflow) (int64, error) {
	var args []interface{}
	args = append(args, workflow.name)
	args = append(args, workflow.pipelineID)
	test := []interface{}{workflow.name, workflow.pipelineID}
	params := execParams{
		insertCmd: "INSERT INTO workflow(name, pipeline_id) VALUES($1, $2) RETURNING id",
		values:    test,
	}

	return p.Client.create(ctx, &params)
}

// createJob will create the provided job and backfill data
// like the ID, CreatedAt, and UpdatedAt fields.
func (p *PipelineService) createJob(ctx context.Context, job *job) (int64, error) {
	var args []interface{}
	args = append(args, job.name)
	args = append(args, job.workflowID)
	args = append(args, job.needs)
	args = append(args, job.steps)
	args = append(args, job.env)
	args = append(args, job.branches)
	args = append(args, job.status)
	params := execParams{
		insertCmd: "INSERT INTO job(name, workflow_id, needs, steps, env, branches, status" +
			") VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING id",
		values: args,
	}

	return p.Client.create(ctx, &params)
}

// createDocker will create the provided docker and backfill data
// like the ID, CreatedAt, and UpdatedAt fields.
func (p *PipelineService) createDocker(ctx context.Context, docker *docker) (int64, error) {
	var args []interface{}
	args = append(args, docker.jobID)
	args = append(args, docker.image)
	args = append(args, docker.tags)
	params := execParams{
		insertCmd: "INSERT INTO job_docker(job_id, image, tags) VALUES($1, $2, $3) RETURNING id",
		values:    args,
	}

	return p.Client.create(ctx, &params)
}

// createMachine will create the provided machine and backfill data
// like the ID, CreatedAt, and UpdatedAt fields.
func (p *PipelineService) createMachine(ctx context.Context, machine *machine) (int64, error) {
	var args []interface{}
	args = append(args, machine.os)
	args = append(args, machine.jobID)
	args = append(args, machine.cpus)
	args = append(args, machine.memory)
	params := execParams{
		insertCmd: "INSERT INTO job_machine(os, job_id, cpus, memory) VALUES($1, $2, $3, $4) RETURNING id",
		values:    args,
	}

	return p.Client.create(ctx, &params)
}
