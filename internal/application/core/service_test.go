package core

import (
	"github.com/stretchr/testify/require"
	"github.com/weezqyd/jpay/internal/adapters/database"
	"github.com/weezqyd/jpay/internal/ports"
	"log"
	"testing"
)

var dbAdapter ports.DbPort
var pagination = 10

func init() {
	var err error
	dbAdapter, err = database.NewAdapter("sqlite3", "./../../../testdb/sample.db", uint64(pagination))
	if err != nil {
		log.Fatalf("could not open db connection err: %v", err)
	}
}
func TestService_GetAllCustomers(t *testing.T) {
	service := NewService(dbAdapter)
	
	data, err := service.GetAllCustomers(1)
	if err != nil {
		log.Fatalf("failed to get all customers err:%v", err)
	}
	//We expect that the data to be paginated based on the provided limit
	require.Equal(t, len(data), pagination)
}

func TestService_GetCustomersByCountry(t *testing.T) {
	service := NewService(dbAdapter)
	
	data, err := service.GetCustomersByCountry("Uganda")
	if err != nil {
		log.Fatalf("failed to get all customers err:%v", err)
	}
	//We expect that the data customers based on country filter
	require.Equal(t, len(data), 3)
	
	data, err = service.GetCustomersByCountry("Unknown")
	if err != nil {
		log.Fatalf("failed to get all customers err:%v", err)
	}
	require.Equal(t, len(data), 14)
}
