package postgres

import (
	"context"

	"github.com/pkg/errors"
	"github.com/souleb/buildeploy/app"
)

// Ensure PipelineService implements app.PipelineService.
var _ app.PipelineService = (*PipelineService)(nil)

// PipelineService represents a PostgreSQL implementation of app.PipelineService.
type PipelineService struct {
	Client *Client
}

// ReadWorkflow will get a workflow by ID.
func (c *Client) ReadWorkflow(ctx context.Context, id uint, workflow *app.Workflow) error {
	err := c.DB.GetContext(ctx, &workflow, "SELECT * FROM workflow WHERE id == $1", id)
	if err != nil {
		return errors.Wrap(err, "sql: ID provided was invalid")
	}

	return nil
}

// GetJobByID will look up a pipeline with the provided ID.
// If the Job is found, we will return a nil error
// If there is an error, we will return an error with
// more information about what went wrong.
func (p *PipelineService) GetJobByID(ctx context.Context, id int64) (*app.Job, error) {
	params := queryParams{
		query: "SELECT * FROM job WHERE id == $1",
		id:    id,
		value: app.Job{},
	}
	err := p.Client.ReadByID(ctx, &params)
	if err != nil {
		return nil, errors.Wrap(err, "failed retrieving the pipeline")
	}

	res := params.value.(*app.Job)
	return res, nil
}

// GetPipelineByID will look up a pipeline with the provided ID.
// If the pipeline is found, we will return a nil error
// If there is an error, we will return an error with
// more information about what went wrong.
func (p *PipelineService) GetPipelineByID(ctx context.Context, id int64) (*app.Pipeline, error) {
	params := queryParams{
		query: "SELECT * FROM pipeline WHERE id == $1",
		id:    id,
		value: app.Pipeline{},
	}
	err := p.Client.ReadByID(ctx, &params)
	if err != nil {
		return nil, errors.Wrap(err, "failed retrieving the pipeline")
	}

	res := params.value.(*app.Pipeline)
	return res, nil
}

// CreatePipeline will create the provided pipeline and backfill data
// like the ID, CreatedAt, and UpdatedAt fields.
func (p *PipelineService) CreatePipeline(ctx context.Context, pipeline *app.Pipeline) (int64, error) {
	var tmp []interface{}
	tmp = append(tmp, pipeline.Name)
	tmp = append(tmp, pipeline.Status)
	tmp = append(tmp, pipeline.CreatedAt)
	tmp = append(tmp, pipeline.DeletedAt)
	params := execParams{
		insertCmd: "INSERT INTO pipeline(name, status, created_at, deleted_at) VALUES($1, $2, $3, $4)",
		value:     tmp,
	}

	return p.Client.Create(ctx, &params)
}

// CreateWorkflow will create the provided workflow and backfill data
// like the ID, CreatedAt, and UpdatedAt fields.
func (p *PipelineService) CreateWorkflow(ctx context.Context, workflow *app.Workflow) (int64, error) {
	var tmp []interface{}
	tmp = append(tmp, workflow.Name)
	tmp = append(tmp, workflow.Name)
	params := execParams{
		insertCmd: "INSERT INTO workflow(name, pipeline_id) VALUES($1, $2)",
		value:     tmp,
	}

	return p.Client.Create(ctx, &params)
}
