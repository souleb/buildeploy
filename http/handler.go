package http

import (
	"context"
	"log"

	"github.com/pkg/errors"
	"github.com/souleb/buildeploy/app"
	pb "github.com/souleb/buildeploy/proto/pipeline/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PipelineHandler struct {
	schemaService    app.SchemaService
	schedulerService app.SchedulerService
	pipelineService  app.PipelineService

	logger *log.Logger
}

func (p *PipelineHandler) store(ctx context.Context, pipeline *app.Pipeline) error {
	pID, err := p.pipelineService.CreatePipeline(ctx, pipeline)
	if err != nil {
		return errors.Wrap(err, "Could not create the pipeline.")
	}

	pipeline.ID = pID

	for k, w := range pipeline.Workflows {
		wID, err := p.pipelineService.CreateWorkflow(ctx, &workflow{name: w.Name, pipelineID: pipeline.ID})
		if err != nil {
			return errors.Wrap(err, "Could not create the pipeline's workflow")
		}
		pipeline.Workflows[k].ID = wID

		for k, j := range w.Jobs {
			jID, err := p.pipelineService.CreateJob(ctx, &job{name: j.Name, workflowID: wID, needs: j.Needs, steps: array(j.Steps),
				env: j.Env, branches: j.Branches, status: j.Status})
			if err != nil {
				return errors.Wrap(err, "Could not create the workflow's job")
			}
			w.Jobs[k].ID = jID
			switch j.Runner.(type) {
			case *app.Docker:
				dID, err := p.pipelineService.CreateDocker(ctx, &docker{image: j.Runner.(*app.Docker).Image, tags: j.Runner.(*app.Docker).Tags, jobID: jID})
				if err != nil {
					return errors.Wrap(err, "Could not create the job's docker runner")
				}
				j.Runner.(*app.Docker).ID = dID
			case *app.Machine:
				mID, err := p.pipelineService.CreateMachine(ctx, &machine{os: j.Runner.(*app.Machine).OS, cpus: j.Runner.(*app.Machine).Cpus,
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

func (p *PipelineHandler) CreatePipeline(ctx context.Context, createPipelineRequest *pb.CreatePipelineRequest) (*pb.CreatePipelineResponse, error) {
	if createPipelineRequest.Item.Name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "`Pipeline` name is not set")
	}

	pipeline := convertToPipeline(createPipelineRequest.GetItem())
	err := p.store(ctx, pipeline)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Could not create the pipeline", err)
	}

	return &pb.CreatePipelineResponse{Id: string(pipeline.ID)}, nil

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
	for _, pbWorkflow := range data.Workflows {
		workflow := convertToWorkflow(pbWorkflow)
		workflows = append(workflows, workflow)
	}

	return &app.Pipeline{
		Name:      data.Name,
		Workflows: workflows,
	}
}

func convertToWorkflow(pbWorkflow *pb.Workflow) app.Workflow {
	jobs := make([]app.Job, 0, len(pbWorkflow.Jobs))
	for _, pbJob := range pbWorkflow.Jobs {
		job := convertToJob(pbJob)
		jobs = append(jobs, job)
	}

	workflow := app.Workflow{
		Name: pbWorkflow.Name,
		Jobs: jobs,
	}

	return workflow

}

func convertToJob(pbJob *pb.Job) app.Job {
	var runnerInstance app.Runner
	switch pbJob.Runner.Type.(type) {
	case *pb.Job_Runner_Docker:
		runnerInstance = &app.Docker{
			Image: pbJob.Runner.Type.(*pb.Job_Runner_Docker).Docker.Image,
			Tags:  pbJob.Runner.Type.(*pb.Job_Runner_Docker).Docker.Tags,
		}
	case *pb.Job_Runner_Machine:
		runnerInstance = &app.Machine{
			OS:     pbJob.Runner.Type.(*pb.Job_Runner_Machine).Machine.Os,
			Cpus:   pbJob.Runner.Type.(*pb.Job_Runner_Machine).Machine.Cpus,
			Memory: pbJob.Runner.Type.(*pb.Job_Runner_Machine).Machine.Memory,
		}
	}

	job := app.Job{
		Name:     pbJob.Name,
		Runner:   runnerInstance,
		Env:      pbJob.Env,
		Steps:    pbJob.Steps.Command,
		Branches: pbJob.Branches,
		Needs:    pbJob.Needs,
	}
	return job
}
