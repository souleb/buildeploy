package app

type Workflow struct {
	ID   int
	Name string
	Job  []Job
}

type WorkflowService interface {
	Workflow(id int) (*Workflow, error)
	CreateWorkflow(w *Workflow) error
}

type Job struct {
	Name   string
	Runner `yaml:",inline"`
	Env    string `yaml:",omitempty"`
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

type Runner struct {
	Kind string   `yaml:"docker"`
	Tags []string `yaml:",omitempty"`
}

type RunnerService interface {
}
