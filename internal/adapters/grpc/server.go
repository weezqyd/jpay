package grpc

import (
	"github.com/weezqyd/jpay/internal/adapters/grpc/pb"
	"github.com/weezqyd/jpay/internal/application/core"
	"google.golang.org/grpc"
	"log"
	"net"
)

// Adapter implements the GRPCPort interface
type Adapter struct {
	core *core.Service
}

// NewAdapter creates a new Adapter
func NewAdapter(api *core.Service) *Adapter {
	return &Adapter{core: api}
}

// Run registers the Customer service to a grpcServer and serves on
// the specified port
func (a Adapter) Run(port string) {
	var err error
	
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen on port %s %v", port, err)
	}
	sever := a
	grpcServer := grpc.NewServer()
	pb.RegisterCustomerServicesServer(grpcServer, sever)
	
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve gRPC server over port 9000: %v", err)
	}
}
