	package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/huynhthanhthao/hrm/proto/user"
	"google.golang.org/grpc"
)

type userServer struct {
	user.UnimplementedUserServiceServer
	users map[string]*user.CreateUserResponse
}

func (s *userServer) CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	id := fmt.Sprintf("user-%d", len(s.users)+1)
	newUser := &user.CreateUserResponse{
		Id:    id,
		Name:  req.Name,
		Email: req.Email,
	}
	s.users[id] = newUser
	return newUser, nil
}


func (s *userServer) GetUser(ctx context.Context, req *user.GetUserRequest) (*user.GetUserResponse, error) {
	foundUser, exists := s.users[req.Id]
	if !exists {
		return nil, fmt.Errorf("user not found")
	}
	return &user.GetUserResponse{
		Id:    foundUser.Id,
		Name:  foundUser.Name,
		Email: foundUser.Email,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server := grpc.NewServer()
	user.RegisterUserServiceServer(server, &userServer{users: make(map[string]*user.CreateUserResponse)})
	log.Println("User Service running on :50051")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}