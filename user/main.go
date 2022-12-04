package main

import (
	"flightBooking/common/config"
	"flightBooking/common/proto"
	"flightBooking/user/dao"
	"flightBooking/user/handler"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	config.Setup()

	dbConfig := config.GetDatabaseConnection()
	userService := dao.NewUserService(dbConfig)

	grpcAddr := fmt.Sprintf("0.0.0.0:%v", config.GetUserGrpcPort())

	// init grpc
	c := make(chan bool)
	go func() {
		list, err := net.Listen("tcp", grpcAddr)
		if err != nil {
			log.Fatalf("Failed to start listener %v", err)
		}
		s := grpc.NewServer()
		server := handler.NewServer(&userService)
		proto.RegisterUserServiceServer(s, server)
		log.Printf("Listening grpc on %v\n", grpcAddr)
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
	case _ = <-time.After(3 * time.Second):
		log.Println("Serving grpc for user-service...")
	}

	// init restful
	r := gin.Default()
	h := handler.NewHandler(&userService)
	r.GET("/ping", h.HandleHealthCheck)

	r.POST("/user", h.HandleRegister)
	r.GET("/user/:username", h.HandleQuery)
	r.PUT("/user/:username", h.HandleUpdate)
	r.POST("/user/validatePassword", h.ValidatePassword)
	err := r.Run("0.0.0.0:8080")
	if err != nil {
		panic(err)
	}
}
