package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/huynhthanhthao/hrm/proto/project"
	"google.golang.org/grpc"
)

type projectServer struct {
	project.UnimplementedProjectServiceServer
	projects map[string]*project.CreateProjectResponse
}

func (s *projectServer) CreateProject(ctx context.Context, req *project.CreateProjectRequest) (*project.CreateProjectResponse, error) {
	id := fmt.Sprintf("project-%d", len(s.projects)+1)
	project := &project.CreateProjectResponse{
		Id:          id,
		Name:        req.Name,
		Description: req.Description,
	}
	s.projects[id] = project
	return project, nil
}

func (s *projectServer) GetProject(ctx context.Context, req *project.GetProjectRequest) (*project.GetProjectResponse, error) {
	projectFound, exists := s.projects[req.Id]
	if !exists {
		return nil, fmt.Errorf("project not found")
	}
	return &project.GetProjectResponse{
		Id:          projectFound.Id,
		Name:        projectFound.Name,
		Description: projectFound.Description,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server := grpc.NewServer()
	project.RegisterProjectServiceServer(server, &projectServer{projects: make(map[string]*project.CreateProjectResponse)})
	log.Println("Project Service running on :50053")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}