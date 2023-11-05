package main

import (
	course_management "course-management-service/coursemanagement"
	"course-management-service/services"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func main() {
	// Create a TCP listener on port 50051
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create a new gRPC server
	s := grpc.NewServer()
	services := &services.Server{}

	// Attach the CourseManagement service to the server
	course_management.RegisterCourseManagementServiceServer(s, services)

	// Start the server
	log.Println("Server is listening on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
