package transport

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/souleb/buildeploy/app"
	pb "github.com/souleb/buildeploy/proto/pipeline/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PipelineHandler struct {
	schemaService   app.SchemaService
	pipelineService app.PipelineService
	m               *sync.Mutex
	pipelines       []*app.Pipeline
}

func (p *PipelineHandler) Fetch() ([]*app.Pipeline, time.Time, error) {
	p.m.Lock()
	defer p.m.Unlock()

	if len(p.pipelines) == 0 {
		return nil, time.Now().Add(1 * time.Millisecond), fmt.Errorf("No pipeline to fetch.")
	}

	response := p.pipelines
	p.pipelines = nil

	return response, time.Now().Add(1 * time.Millisecond), nil
}

func (p *PipelineHandler) store(pipeline *app.Pipeline) {
	p.m.Lock()
	defer p.m.Unlock()
	p.pipelines = append(p.pipelines, pipeline)
}

func (p *PipelineHandler) CreatePipeline(ctx context.Context, createPipelineRequest *pb.CreatePipelineRequest) (*pb.CreatePipelineResponse, error) {
	if createPipelineRequest.Item.Name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "`Pipeline` name is not set")
	}

	pipeline := convertToPipeline(createPipelineRequest.GetItem())
	err := p.pipelineService.CreatePipeline(ctx, pipeline)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Could not create the pipeline", err)
	}

	go p.store(pipeline)
	/*
		err = p.schedulerService.Schedule(pipeline)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Could not run the Pipeline.", err)
		}*/

	return &pb.CreatePipelineResponse{Id: fmt.Sprint(pipeline.ID)}, nil

	/*err := wh.SchemaService.Validate(w)
	if err != nil {
		return nil, err
	}*/

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
