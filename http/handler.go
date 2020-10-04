package http

//
//import (
//	"fmt"
//	"net/http"
//	"strings"
//)
//
//type Handler struct {
//	WorkflowHandler *WorkflowHandler
//}
//
//func NewHandler() *Handler {
//	return &Handler{
//		WorkflowHandler: NewWorkflowHandler(),
//	}
//}
//
//// ServeHTTP delegates a request to the appropriate subhandler.
//func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	if strings.HasPrefix(r.URL.Path, "/api/workflow") {
//		h.WorkflowHandler.Router.ServeHTTP(w, r)
//	} else {
//		fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
//	}
//}
