package http

// TO DO:
// Rename to delivery

import (
	"net"

	"github.com/souleb/buildeploy/app"
	pb "github.com/souleb/buildeploy/proto/pipeline/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const defaultAdrr = ":3000"

type Server struct {
	ln net.Listener
	//Handler    *Handler
	grpcServer       *grpc.Server
	Addr             string
	schedulerService app.SchedulerService
	pipelineService  app.PipelineService
}

func NewServer(schedulerService app.SchedulerService, pipelineService app.PipelineService) (*Server, error) {

	//creds, _ := credentials.NewServerTLSFromFile(certFile, keyFile)
	server := &Server{
		//Handler:    NewHandler(),
		//grpcServer: grpc.NewServer(grpc.Creds(creds)),
		grpcServer:       grpc.NewServer(),
		Addr:             defaultAdrr,
		schedulerService: schedulerService,
		pipelineService:  pipelineService,
	}

	ln, err := net.Listen("tcp", server.Addr)
	if err != nil {
		return nil, err
	}
	server.ln = ln

	return server, nil
}

func (s *Server) Open() error {

	pb.RegisterPipelineServiceServer(s.grpcServer, &PipelineHandler{SchedulerService: s.schedulerService, PipelineService: s.pipelineService})
	// Register reflection service on gRPC server.
	reflection.Register(s.grpcServer)
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
