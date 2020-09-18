package postgres

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/souleb/buildeploy/app"
)

// Ensure WorkflowService implements app.WorkflowService.
var _ app.WorkflowService = &WorkflowService{}

// WorkflowService represents a PostgreSQL implementation of app.WorkflowService.
type WorkflowService struct {
	Client *Client
}

// ByID will look up a workflow with the provided ID.
// If the workflow is found, we will return a nil error
// If there is an error, we will return an error with
// more information about what went wrong.
func (w *WorkflowService) ByID(id int) (*app.Workflow, error) {
	var workflow app.Workflow
	err := w.Client.DB.Where("id = ?", id).First(&workflow).Error
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("gorm: reading workflow with %d failed", id))
	}

	return &workflow, nil
}

// Create will create the provided workflow and backfill data
// like the ID, CreatedAt, and UpdatedAt fields.
func (w *WorkflowService) Create(workflow *app.Workflow) error {
	return w.Client.DB.Create(workflow).Error
}

// DestructiveReset drops the workflow table and rebuilds it
//func (w *WorkflowService) DestructiveReset() {
////	w.Client.AutoMigrate(&app.Workflow{})
//}
