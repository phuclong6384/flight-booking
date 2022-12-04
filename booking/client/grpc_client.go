package client

import (
	"context"
	"flightBooking/common/config"
	"flightBooking/common/proto"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcClient struct {
	userClient   proto.UserServiceClient
	flightClient proto.FlightServiceClient
}

func NewUserGrpcClient() GrpcClient {
	userGrpcAddr := fmt.Sprintf("0.0.0.0:%v", config.GetUserGrpcPort())
	userConn, err := grpc.Dial(userGrpcAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error dial %v\n", err)
	}

	flightGrpcAddr := fmt.Sprintf("0.0.0.0:%v", config.GetFlightGrpcPort())
	flightConn, err := grpc.Dial(flightGrpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error dial %v\n", err)
	}
	return GrpcClient{
		userClient:   proto.NewUserServiceClient(userConn),
		flightClient: proto.NewFlightServiceClient(flightConn),
	}
}

func (g *GrpcClient) ValidatePasswordUserService(username string, password string) bool {
	validatePassword, err := g.userClient.ValidatePassword(context.Background(),
		&proto.ValidatePasswordRequest{Username: username, Password: password})
	if err != nil {
		log.Println("Error occurred when calling grpc request to user")
		return false
	}
	return validatePassword.Success
}

func (g *GrpcClient) GetFlightDetail(id int) (*proto.FlightDetailResponse, error) {
	flight, err := g.flightClient.SearchFlight(context.Background(), &proto.QueryFlightDetail{FlightId: int32(id)})
	if err != nil {
		log.Println("Error occurred when calling grpc request to flight")
		return nil, err
	}
	return flight, nil
}
