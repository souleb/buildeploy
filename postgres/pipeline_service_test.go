package postgres

import (
	"testing"
	"time"

	"github.com/souleb/buildeploy/app"
	"golang.org/x/net/context"
)

func TestCreatePipeline(t *testing.T) {
	service := PipelineService{
		Client: testClient,
	}

	pipeline := app.Pipeline{
		Name:      "TestCreatePipeline",
		Status:    0,
		CreatedAt: time.Now(),
	}

	workflow := workflow{
		name: "TestCreateWorkflow",
	}

	job := job{
		name:     "TestCreateWorkflow",
		edges:    "job1, job2",
		steps:    []string{"mkdir test", "rm -Rf test"},
		env:      "develop",
		branches: "feature*, develop",
		status:   0,
	}

	ctx := context.Background()

	t.Run("createPipeline", func(t *testing.T) {
		id, err := service.CreatePipeline(ctx, &pipeline)

		if err != nil {
			t.Errorf("TestCreatePipeline failed while creating the pipeline %s", err)
		}

		if id == 0 {
			t.Errorf("TestCreatePipeline received unexpected id %d while creating the pipeline", id)
		}

		pipeline.ID = id
	})

	t.Run("createWorkflow", func(t *testing.T) {
		workflow.pipelineID = pipeline.ID

		id, err := service.createWorkflow(context.Background(), &workflow)

		if err != nil {
			t.Errorf("TestCreateWorkflow failed while creating the workflow %s", err)
		}

		if id == 0 {
			t.Errorf("TestCreateWorkflow received unexpected id %d while creating the workflow", id)
		}

		workflow.id = id
	})

	t.Run("createJob", func(t *testing.T) {
		job.workflowID = workflow.id

		id, err := service.createJob(context.Background(), &job)

		if err != nil {
			t.Errorf("TestCreateJob failed while creating the job %s", err)
		}

		if id == 0 {
			t.Errorf("TestCreateJob received unexpected id %d while creating the job", id)
		}

		job.id = id
	})
}
