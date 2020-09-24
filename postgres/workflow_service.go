package postgres

import (
	"github.com/pkg/errors"
	"github.com/souleb/buildeploy/app"
)

// Ensure WorkflowService implements app.WorkflowService.
var _ app.WorkflowService = &WorkflowService{}

// WorkflowService represents a PostgreSQL implementation of app.WorkflowService.
type WorkflowService struct {
	Client *Client
}

// GetByID will look up a workflow with the provided ID.
// If the workflow is found, we will return a nil error
// If there is an error, we will return an error with
// more information about what went wrong.
func (w *WorkflowService) GetByID(id uint) (*app.Workflow, error) {
	value, err := w.Client.Read(id)
	if err != nil {
		return nil, errors.Wrap(err, "failed retrieving workflow")
	}

	return value, nil
}

// Create will create the provided workflow and backfill data
// like the ID, CreatedAt, and UpdatedAt fields.
func (w *WorkflowService) Create(workflow *app.Workflow) error {
	return w.Client.CreateWorkflow(workflow)
}

// DestructiveReset drops the workflow table and rebuilds it
func (w *WorkflowService) DestructiveReset() {
	//w.Client.DB.Migrator().DropTable(&app.Workflow{})
	w.Client.DB.AutoMigrate(&app.Workflow{})
	w.Client.DB.AutoMigrate(&app.Job{})
}
