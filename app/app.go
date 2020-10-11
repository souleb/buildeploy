package app

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"strings"
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
	ID         uint
	WorkflowID uint
	Status     Status
}

// Workflow defines a set of jobs constituing the workflow.
type Workflow struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Jobs      []Job
	Status    Status
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}

// WorkflowService represents a service for managing workflows.
type WorkflowService interface {
	GetByID(id uint) (*Workflow, error)
	Create(w *Workflow) error
}

// Job is a defined set of steps to execute
// It uses a defined executor to do so
type Job struct {
	ID         uint `gorm:"primaryKey"`
	Name       string
	Runner     Runner
	Steps      Commands
	Env        string
	Branches   string
	Needs      []string
	Status     Status
	WorkflowID uint
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time `gorm:"index"`
}

// Commands are actions to be performed
type Commands []string

func (commands Commands) Value() (driver.Value, error) {
	var quotedStrings []string
	for _, str := range commands {
		quotedStrings = append(quotedStrings, strconv.Quote(str))
	}

	value := fmt.Sprintf("{ %s }", strings.Join(quotedStrings, ","))

	return value, nil
}

func (commands *Commands) Scan(src interface{}) error {
	val, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("unable to scan")
	}
	value := strings.TrimPrefix(string(val), "{")
	value = strings.TrimSuffix(value, "}")

	*commands = strings.Split(value, ",")

	return nil
}

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
	OS       string
	CPUCores string
	Memory   string
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
