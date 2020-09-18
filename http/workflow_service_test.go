package http

import (
	"context"
	"testing"

	pb "github.com/souleb/buildeploy/proto/workflow/v1"
)

func CreateWorkflowTest(t *testing.T) {
	tests := []struct {
		input          pb.Workflow
		expectedOutput string
	}{
		{
			pb.Workflow{
				Id:   1,
				Name: "test1",
				Jobs: []*pb.Job{
					{
						Name: "job1",
						Executor: &pb.Job_Docker{
							Docker: &pb.Docker{
								Image: "docker.io/image-test",
								Tags:  "stable, golang",
							},
						},
						Steps: []*pb.Step{
							{
								StepCommand: &pb.Step_Command{
									Command: &pb.Command{
										Command: "ls -l $PWD",
									},
								},
							},
						},
						Env:      "integration",
						Branches: "develop",
					},
				},
			},
			"testid",
		},
		{
			pb.Workflow{
				Id:   2,
				Name: "test2",
				Jobs: []*pb.Job{
					{
						Name: "job1",
						Executor: &pb.Job_Docker{
							Docker: &pb.Docker{
								Image: "docker.io/image-test2",
								Tags:  "stable, golang",
							},
						},
						Steps: []*pb.Step{
							{
								StepCommand: &pb.Step_Command{
									Command: &pb.Command{
										Command: "curl www.google.com",
									},
								},
							},
						},
						Env:      "release",
						Branches: "master, release",
					},
				},
			},
			"testid",
		},
	}

	w := &WorkflowHandler{}

	for _, tt := range tests {
		req := &pb.CreateWorkflowRequest{
			Item: &tt.input,
		}

		resp, err := w.CreateWorkflow(context.Background(), req)
		if err != nil {
			t.Errorf("CreateWorkflowTest(%v) got unexpected error", err)
		}

		if resp.GetId() != tt.expectedOutput {
			t.Errorf("CreateWorkflowTest(%v)=%v, wanted %v", tt.input.Name, resp.GetId(), tt.expectedOutput)
		}
	}

}
