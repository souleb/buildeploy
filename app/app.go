package app

import (
	"context"
	"time"
)

// Keyword defines a built-in command.
// It will be interpreted by the application.
// And perform and pre-defined task based on the actual context.
type Keyword string

type Status int

const (
	CHECKOUT Keyword = "checkout"
)

const (
	FAILED Status = iota - 1
	BLOCKED
	QUEUED
	RUNNING
	SUCCESS
)

type Pipeline struct {
	ID   int64
	Name string
	//WorkflowID int64
	Workflows []Workflow
	Status    Status
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

// Workflow defines a set of jobs constituing the workflow.
type Workflow struct {
	ID   int64
	Name string
	Jobs []Job
}

// PipelineService represents a service for managing pipelines.
type PipelineService interface {
	GetJobByID(ctx context.Context, id int64) (*Job, error)
	GetPipelineByID(ctx context.Context, id int64) (*Pipeline, error)
	CreatePipeline(ctx context.Context, pipeline *Pipeline) (int64, error)
}

// Job is a defined set of steps to execute
// It uses a defined executor to do so
type Job struct {
	ID       int64
	Name     string
	Runner   Runner
	Steps    Commands
	Env      string
	Branches string
	Needs    []string
	Status   Status
	//WorkflowID uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// Commands are actions to be performed
type Commands []string

type Runner interface {
	isJobRunner()
}

// Docker is an Executor type
type Docker struct {
	Image string
	Tags  string
}

func (d *Docker) isJobRunner() {}

// Machine is an executor type
type Machine struct {
	OS     string
	Cpus   string
	Memory string
}

func (m *Machine) isJobRunner() {}

// SchemaService represents a service for managing jsonschemas.
type SchemaService interface {
	Validate(data interface{}) error
}

// SchedulerService represents a service for managing schedulers.
type SchedulerService interface {
	Schedule(workflow *Workflow) error
}
