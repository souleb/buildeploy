package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/souleb/buildeploy/app"
	pb "github.com/souleb/buildeploy/proto/workflow/v1"
)

type WorkflowHandler struct {
	*httprouter.Router

	WorkflowService app.WorkflowService

	Logger *log.Logger
}

func NewWorkflowHandler() *WorkflowHandler {
	h := &WorkflowHandler{

		Router: httprouter.New(),
		Logger: log.New(os.Stderr, "", log.LstdFlags),
	}

	//h.Router.POST("/api/workflows", h.handlePostWorkflow)
	h.Router.GET("/api/workflow/:id", h.handleGetWorkflow)
	h.Router.GET("/api/workflow/", h.handleGetWorkflows)
	//h.Router.PATCH("/api/workflow/:id", h.handlePatchWorkflow)

	return h

}

func (wh *WorkflowHandler) handleGetWorkflows(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprint(w, "<h1>Welcome to my handleGetWorkflow SSSSS!</h1>")
}

func (wh *WorkflowHandler) handleGetWorkflow(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprint(w, "<h1>Welcome to my handleGetWorkflow!</h1>")
}

func (wh *WorkflowHandler) CreateWorkflow(ctx context.Context, createWorkflowRequest *pb.CreateWorkflowRequest) (*pb.CreateWorkflowResponse, error) {
	return &pb.CreateWorkflowResponse{Id: "testid"}, nil
}
