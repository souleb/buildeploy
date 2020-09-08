package http

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/SouleBA/buildeploy/app"
	"github.com/julienschmidt/httprouter"
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
	/*id := ps.ByName("id")

	// Find dial by ID.
	d, err := wh.WorkflowService.Workflow(app.Workflow(id))
	if err != nil {
		fmt.Fprintf(w, err, http.StatusInternalServerError, wh.Logger)
	} else if d == nil {
		NotFound(w)
	} else {
		d, err = json.Marshal(w, &getDialResponse{Dial: d})
		, h.Logger)
	}*/
	fmt.Fprint(w, "<h1>Welcome to my handleGetWorkflow!</h1>")
}

type getDialResponse struct {
	Dial *app.Workflow `json:"dial,omitempty"`
	Err  string        `json:"err,omitempty"`
}

// ServeHTTP delegates a request to the appropriate subhandler.
/*func (wh *WorkflowHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "<h1>Welcome to my awesome site workflow handler!</h1>")

}*/
