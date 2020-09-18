package http

import (
	"net"

	pb "github.com/souleb/buildeploy/proto/workflow/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// DefaultAddr is the default bind address.
const defaultAdrr = ":3000"

type Server struct {
	ln net.Listener
	//Handler    *Handler
	grpcServer *grpc.Server
	Addr       string
}

func NewServer() (*Server, error) {

	//creds, _ := credentials.NewServerTLSFromFile(certFile, keyFile)
	server := &Server{
		//Handler:    NewHandler(),
		//grpcServer: grpc.NewServer(grpc.Creds(creds)),
		grpcServer: grpc.NewServer(),
		Addr:       defaultAdrr,
	}

	ln, err := net.Listen("tcp", server.Addr)
	if err != nil {
		return nil, err
	}
	server.ln = ln

	return server, nil
}

func (s *Server) Open() error {

	pb.RegisterWorkflowServiceServer(s.grpcServer, &WorkflowHandler{})
	// Register reflection service on gRPC server.
	reflection.Register(s.grpcServer)
	// Start HTTP server.
	s.grpcServer.Serve(s.ln)

	return nil
}

// Close closes the socket.
func (s *Server) Close() error {
	s.grpcServer.Stop()

	if s.ln != nil {
		s.ln.Close()
	}
	return nil
}
