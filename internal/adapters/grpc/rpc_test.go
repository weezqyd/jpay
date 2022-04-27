package grpc

import (
	"context"
	"github.com/stretchr/testify/require"
	"github.com/weezqyd/jpay/internal/adapters/database"
	"github.com/weezqyd/jpay/internal/adapters/grpc/pb"
	"github.com/weezqyd/jpay/internal/application/core"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
	"testing"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	var err error
	lis = bufconn.Listen(bufSize)
	grpcServer := grpc.NewServer()
	dbAdapter, err := database.NewAdapter("sqlite3", "./../../../testdb/sample.db", 10)
	if err != nil {
		log.Fatalf("failed to initiate dbase connection: %v", err)
	}
	
	// core
	service := core.NewService(dbAdapter)
	gRPCAdapter := NewAdapter(service)
	
	pb.RegisterCustomerServicesServer(grpcServer, gRPCAdapter)
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("test server start error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func getGRPCConnection(ctx context.Context, t *testing.T) *grpc.ClientConn {
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial bufnet: %v", err)
	}
	return conn
}
func TestAdapterGetAllCustomers(t *testing.T) {
	ctx := context.Background()
	conn := getGRPCConnection(ctx, t)
	defer conn.Close()
	
	client := pb.NewCustomerServicesClient(conn)
	country := "Uganda"
	
	params := &pb.GetCustomersOptions{
		Country: &country,
	}
	
	answer, err := client.GetAllCustomers(ctx, params)
	if err != nil {
		t.Fatalf("expected: %v, got: %v", nil, err)
	}
	
	require.Equal(t, len(answer.Data), 3)
}
