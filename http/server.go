package http

import (
	"net"
	"net/http"

	pb "github.com/souleb/buildeploy/proto/workflow/v1"
	"google.golang.org/grpc"
)

// DefaultAddr is the default bind address.
const defaultAdrr = ":3000"

type Server struct {
	ln         net.Listener
	Handler    *Handler
	grpcServer *grpc.Server
	Addr       string
}

func NewServer() *Server {
	return &Server{
		//Handler:    NewHandler(),
		grpcServer: grpc.NewServer(),
		Addr:       defaultAdrr,
	}
}

func (s *Server) Open() error {
	ln, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return err
	}

	pb.CreateWorkflowRequest(s.grpcServer, &WorkflowHandler{})
	//s.ln = ln
	// Start HTTP server.
	go func() { http.Serve(s.ln, s.Handler) }()

	return nil
}

// Close closes the socket.
func (s *Server) Close() error {
	if s.ln != nil {
		s.ln.Close()
	}
	return nil
}
