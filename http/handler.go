package http

import (
	"context"
	"log"

	"github.com/pkg/errors"
	"github.com/souleb/buildeploy/app"
	pb "github.com/souleb/buildeploy/proto/workflow/v1"
)

type WorkflowHandler struct {
	SchemaService    app.SchemaService
	SchedulerService app.SchedulerService

	Logger *log.Logger
}

func (wh *WorkflowHandler) CreateWorkflow(ctx context.Context, createWorkflowRequest *pb.CreateWorkflowRequest) (*pb.CreateWorkflowResponse, error) {
	w := convertToWorkflow(createWorkflowRequest.Item)

	/*err := wh.SchemaService.Validate(w)
	if err != nil {
		return nil, err
	}*/

	err := wh.SchedulerService.Schedule(w)
	if err != nil {
		return nil, errors.Wrap(err, "Impossible to schedule the workflow.")
	}

	return &pb.CreateWorkflowResponse{Id: "testid"}, nil
}

func convertToWorkflow(data *pb.Workflow) *app.Workflow {
	jobs := make([]app.Job, 0, len(data.Jobs))
	for _, job := range data.Jobs {
		var runnerInstance app.Runner
		switch job.Runner.Type.(type) {
		case *pb.Job_Runner_Docker:
			runnerInstance = &app.Docker{
				Image: job.Runner.Type.(*pb.Job_Runner_Docker).Docker.Image,
				Tags:  job.Runner.Type.(*pb.Job_Runner_Docker).Docker.Tags,
			}
		case *pb.Job_Runner_Machine:
			runnerInstance = &app.Machine{
				OS:     job.Runner.Type.(*pb.Job_Runner_Machine).Machine.Os,
				Cpus:   job.Runner.Type.(*pb.Job_Runner_Machine).Machine.Cpus,
				Memory: job.Runner.Type.(*pb.Job_Runner_Machine).Machine.Memory,
			}
		}

		newJob := app.Job{
			Name:     job.Name,
			Runner:   runnerInstance,
			Env:      job.Env,
			Steps:    job.Steps.Command,
			Branches: job.Branches,
			Needs:    job.Needs,
		}
		jobs = append(jobs, newJob)
	}

	return &app.Workflow{
		Name: data.Name,
		Jobs: jobs,
	}
}
