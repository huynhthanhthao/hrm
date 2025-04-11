package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"hrm/proto/employee"
	"google.golang.org/grpc"
)

type employeeServer struct {
	employee.UnimplementedEmployeeServiceServer
	employees map[string]*employee.CreateEmployeeResponse
}

func (s *employeeServer) CreateEmployee(ctx context.Context, req *employee.CreateEmployeeRequest) (*employee.CreateEmployeeResponse, error) {
	id := fmt.Sprintf("employee-%d", len(s.employees)+1)
	employee := &employee.CreateEmployeeResponse{
		Id:         id,
		Name:       req.Name,
		Department: req.Department,
	}
	s.employees[id] = employee
	return employee, nil
}

func (s *employeeServer) GetEmployee(ctx context.Context, req *employee.GetEmployeeRequest) (*employee.GetEmployeeResponse, error) {
	employeeFound, exists := s.employees[req.Id]
	if !exists {
		return nil, fmt.Errorf("employee not found")
	}
	return &employee.GetEmployeeResponse{
		Id:         employeeFound.Id,
		Name:       employeeFound.Name,
		Department: employeeFound.Department,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server := grpc.NewServer()
	employee.RegisterEmployeeServiceServer(server, &employeeServer{employees: make(map[string]*employee.CreateEmployeeResponse)})
	log.Println("Employee Service running on :50052")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}