package app

type Workflow struct {
	ID   int
	Name string
	Jobs []Job
}

type WorkflowService interface {
	Workflow(id int) (*Workflow, error)
	CreateWorkflow(w *Workflow) error
}

type Job struct {
	Name     string
	Executor `yaml:",inline"`
	Env      string `yaml:",omitempty"`
	//Steps    []Step   `yaml:",inline"`
	Branches string
}

type JobService interface {
}

type Step struct {
	Checkout string   `yaml:",omitempty"`
	Command  []string `yaml:",omitempty"`
}

type StepService interface {
}

type Executor struct {
	Kind string   `yaml:"docker"`
	Tags []string `yaml:",omitempty"`
}

type RunnerService interface {
}

type SchemaService interface {
	Validate(data interface{})
}
