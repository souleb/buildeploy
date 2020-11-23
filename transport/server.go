package transport

// TO DO:
// Rename to delivery

import (
	"net"
	"time"

	"github.com/souleb/buildeploy/app"
	pb "github.com/souleb/buildeploy/proto/pipeline/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const defaultAdrr = ":3000"

const maxPending = 10

type Server struct {
	ln net.Listener
	//Handler    *Handler
	grpcServer      *grpc.Server
	Addr            string
	pipelineService app.PipelineService
	apiList         map[string]interface{} // keep Apis references to ask them to add watches
}

func NewServer(pipelineService app.PipelineService) (*Server, error) {

	//creds, _ := credentials.NewServerTLSFromFile(certFile, keyFile)
	server := &Server{
		//Handler:    NewHandler(),
		//grpcServer: grpc.NewServer(grpc.Creds(creds)),
		grpcServer:      grpc.NewServer(),
		Addr:            defaultAdrr,
		pipelineService: pipelineService,
	}

	ln, err := net.Listen("tcp", server.Addr)
	if err != nil {
		return nil, err
	}
	server.ln = ln

	return server, nil
}

func (s *Server) Open() error {

	pipelineHandler := PipelineHandler{pipelineService: s.pipelineService}
	s.apiList["pipelineService"] = &pipelineHandler
	pb.RegisterPipelineServiceServer(s.grpcServer, &pipelineHandler)
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

// Subscribe register an event listener to an api event
func (s *Server) Subscribe(api string) app.Subscription {
	switch api {
	case "pipelineService":
		s := &sub{
			api:     s.apiList[api].(*PipelineHandler),
			updates: make(chan *app.Pipeline),
		}
		go s.loop()
		return s
	default:
		return nil
	}
}

type sub struct {
	api     *PipelineHandler
	updates chan *app.Pipeline
	closing chan chan error
}

func (s *sub) loop() {
	var pending []*app.Pipeline
	var next time.Time
	var err error
	for {
		var first *app.Pipeline
		var updates chan *app.Pipeline
		if len(pending) > 0 {
			first = pending[0]
			updates = s.updates
		}
		var fetchDelay time.Duration
		if now := time.Now(); next.After(now) {
			fetchDelay = next.Sub(now)
		}
		var startFetch <-chan time.Time
		if len(pending) < maxPending {
			startFetch = time.After(fetchDelay)
		}
		select {
		case <-startFetch:
			var fetched []*app.Pipeline
			fetched, next, err = s.api.Fetch()
			if err != nil {
				next = time.Now().Add(10 * time.Millisecond)
				break
			}
			pending = append(pending, fetched...)
		case updates <- first:
			pending = pending[1:]
		case errc := <-s.closing:
			errc <- err
			close(s.updates)
		}
	}

}

func (s *sub) Updates() <-chan *app.Pipeline {
	return s.updates
}

func (s *sub) Close() error {
	errc := make(chan error)
	s.closing <- errc
	return <-errc
}
