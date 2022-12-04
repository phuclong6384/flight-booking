package handler

import (
	"context"
	"flightBooking/common/proto"
	"flightBooking/user/dao"
	"log"
)

type Server struct {
	proto.UserServiceServer
	UserService dao.IUserService
}

func NewServer(userService dao.IUserService) proto.UserServiceServer {
	return &Server{
		UserService: userService,
	}
}

func (s *Server) ValidatePassword(ctx context.Context, req *proto.ValidatePasswordRequest) (*proto.ValidatePasswordResponse, error) {
	user, err := s.UserService.GetByUsername(req.Username)
	if err != nil {
		log.Println("Cannot find user", err)
		return &proto.ValidatePasswordResponse{Success: false}, nil
	}
	valid := s.UserService.ValidatePassword(user, req.Password)
	return &proto.ValidatePasswordResponse{Success: valid}, nil
}
