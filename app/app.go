package app

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Keyword string

const (
	CHECKOUT Keyword = "checkout"
)

type Workflow struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Jobs      []Job
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type WorkflowService interface {
	ByID(id int) (*Workflow, error)
	Create(w *Workflow) error
}

type Job struct {
	ID   uint `gorm:"primaryKey"`
	Name string
	Executor
	Steps     []Step `gorm:"embedded"`
	Env       string
	Branches  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type JobService interface {
}

type Command string

type Step struct {
	Keyword
	Command
}

type StepService interface {
}

type Executor interface {
	Kind() string
}

type Docker struct {
	Image string
	Tags  string
}

func (d *Docker) Kind() string {
	return fmt.Sprintf("%T", *d)
}

type Machine struct {
	OS       string
	CpuCores string
	Memory   string
}

func (m *Machine) Kind() string {
	return fmt.Sprintf("%T", *m)
}

type RunnerService interface {
}

type SchemaService interface {
	Validate(data interface{})
}
