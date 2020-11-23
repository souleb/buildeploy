package transport

import (
	"context"
	"testing"

	pb "github.com/souleb/buildeploy/proto/pipeline/v1"
)

func TestCreatePipeline(t *testing.T) {
	tests := []struct {
		name           string
		input          pb.Pipeline
		expectedOutput string
	}{
		{
			name: "testcase1",
			input: pb.Pipeline{
				Name: "pipeline1",
				Workflows: []*pb.Workflow{
					{
						Id:   1,
						Name: "workflow1",
						Jobs: []*pb.Job{
							{
								Name: "job1",
								Runner: &pb.Job_Runner{
									Type: &pb.Job_Runner_Docker{
										Docker: &pb.Docker{
											Image: "docker.io/image-test",
											Tags:  "stable, golang",
										},
									},
								},
								Steps: &pb.Steps{
									Command: []string{
										"ls -l $PWD",
									},
								},
								Env:      "integration",
								Branches: "develop",
							},
						},
					},
				},
			},
			expectedOutput: "testid",
		},
		{
			name: "testcase2",
			input: pb.Pipeline{
				Name: "pipeline2",
				Workflows: []*pb.Workflow{
					{
						Id:   2,
						Name: "workflow2",
						Jobs: []*pb.Job{
							{
								Name: "job1",
								Runner: &pb.Job_Runner{
									Type: &pb.Job_Runner_Docker{
										Docker: &pb.Docker{
											Image: "docker.io/image-test2",
											Tags:  "stable, golang",
										},
									},
								},
								Steps: &pb.Steps{
									Command: []string{
										"curl www.google.com",
									},
								},
								Env:      "release",
								Branches: "master, release",
							},
						},
					},
				},
			},
			expectedOutput: "testid",
		},
	}

	p := &PipelineHandler{}

	for k := range tests {
		t.Run(tests[k].name, func(t *testing.T) {
			req := &pb.CreatePipelineRequest{
				Item: &tests[k].input,
			}
			resp, err := p.CreatePipeline(context.Background(), req)
			if err != nil {
				t.Errorf("CreateWorkfPipelineTest(%v) got unexpected error", err)
			}

			if resp.GetId() != tests[k].expectedOutput {
				t.Errorf("CreatePipelineTest(%v)=%v, wanted %v", tests[k].input.Name, resp.GetId(), tests[k].expectedOutput)
			}
		})
	}

}
