package main

import (
	"flag"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/weezqyd/jpay/internal/adapters/database"
	gRPC "github.com/weezqyd/jpay/internal/adapters/grpc"
	"github.com/weezqyd/jpay/internal/application/api"
	"github.com/weezqyd/jpay/internal/application/core"
	"log"
	"time"
)

func main() {
	var dsn string
	var paginationLimit uint64
	
	flag.StringVar(&dsn, "dsn", "./testdb/sample.db", "Data source name")
	flag.Uint64Var(&paginationLimit, "limit", 15, "Default pagination limit")
	flag.Parse()
	
	db, err := database.NewAdapter("sqlite3", dsn, paginationLimit)
	if err != nil {
		log.Fatalf("could not open database connection")
	}
	defer db.CloseDbConnection()
	service := core.NewService(db)
	
	go func(s *core.Service) {
		log.Println("Starting gRPC server on port 9090")
		server := gRPC.NewAdapter(s)
		server.Run(":9090")
	}(service)
	
	apiPort := api.NewApi(core.NewService(db))
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:3001", "http://localhost:9000"},
		AllowMethods:     []string{"POST", "GET", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	router.Use(static.Serve("/", static.LocalFile("./frontend/dist/", false)))
	router.GET("/api/v1/customers", apiPort.GetAllCustomers)
	log.Fatalln(router.Run(":9000"))
}
