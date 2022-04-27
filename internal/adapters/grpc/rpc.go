package grpc

import (
	"context"
	"github.com/weezqyd/jpay/internal/adapters/grpc/pb"
	"github.com/weezqyd/jpay/internal/application/core"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a Adapter) GetAllCustomers(ctx context.Context, options *pb.GetCustomersOptions) (*pb.CustomersResponse, error) {
	var err error
	var customers []*core.Customer
	if options.GetCountry() != "" {
		customers, err = a.core.GetCustomersByCountry(options.GetCountry())
	} else {
		customers, err = a.core.GetAllCustomers(options.GetPage())
	}
	
	ans := &pb.CustomersResponse{}
	if err != nil {
		return ans, status.Error(codes.Internal, "unexpected error")
	}
	var data []*pb.Customer
	for _, c := range customers {
		customer := &pb.Customer{
			Id:      c.Id,
			Name:    c.Name,
			Phone:   c.Phone,
			Country: c.Country,
		}
		data = append(data, customer)
	}
	ans = &pb.CustomersResponse{
		Data: data,
	}
	
	return ans, nil
}
