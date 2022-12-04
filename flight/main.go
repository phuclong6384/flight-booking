package main

import (
	"flightBooking/common/config"
	"flightBooking/common/proto"
	"flightBooking/flight/dao"
	"flightBooking/flight/handler"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	config.Setup()
	config.GetDatabaseConnection()

	dbConfig := config.GetDatabaseConnection()
	flightService := dao.NewFlightService(dbConfig)

	grpcAddr := fmt.Sprintf("0.0.0.0:%v", config.GetFlightGrpcPort())

	// init grpc
	c := make(chan bool)
	go func() {
		list, err := net.Listen("tcp", grpcAddr)
		if err != nil {
			log.Fatalf("Failed to start listener %v", err)
		}
		log.Printf("Listening on %v\n", grpcAddr)
		s := grpc.NewServer()
		server := &handler.Server{FlightService: &flightService}
		proto.RegisterFlightServiceServer(s, server)
		if err = s.Serve(list); err != nil {
			c <- false
			log.Fatalf("Failed to serve %v\n", err)
		}
	}()
	select {
	case success := <-c:
		if !success {
			panic("Cannot init grpc")
		}
	case _ = <-time.After(1 * time.Second):
		log.Println("Serving grpc for user-service...")
	}

	r := gin.Default()
	h := handler.NewHandler(&flightService)
	r.GET("/ping", h.HandleHealthCheck)

	r.GET("/flight", h.HandleGetList)
	r.GET("/flight/:code", h.HandleGetByCode)
	err := r.Run("0.0.0.0:8081")
	if err != nil {
		panic(err)
	}
}
