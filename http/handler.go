package http

import (
	"context"
	"log"

	"github.com/souleb/buildeploy/app"
	pb "github.com/souleb/buildeploy/proto/pipeline/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PipelineHandler struct {
	SchemaService    app.SchemaService
	SchedulerService app.SchedulerService

	Logger *log.Logger
}

func (wh *PipelineHandler) CreatePipeline(ctx context.Context, createPipelineRequest *pb.CreatePipelineRequest) (*pb.CreatePipelineResponse, error) {
	if createPipelineRequest.Item.Name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "`Pipeline` name is not set")
	}
	//p := convertToPipeline(createPipelineRequest.Item)

	/*err := wh.SchemaService.Validate(w)
	if err != nil {
		return nil, err
	}*/

	//err := wh.SchedulerService.Schedule(p)
	//if err != nil {
	//	return nil, errors.Wrap(err, "Impossible to schedule the Pipeline.")
	//}

	return &pb.CreatePipelineResponse{Id: "testid"}, nil
}

func convertToPipeline(data *pb.Pipeline) *app.Pipeline {
	workflows := make([]app.Workflow, 0, len(data.Workflows))
	for _, wk := range data.Workflows {
		jobs := make([]app.Job, 0, len(wk.Jobs))
		for _, job := range wk.Jobs {
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

		newWorkflow := app.Workflow{
			Name: data.Name,
			Jobs: jobs,
		}

		workflows = append(workflows, newWorkflow)
	}

	return &app.Pipeline{
		Name:      data.Name,
		Workflows: workflows,
	}
}
