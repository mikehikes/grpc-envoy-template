package main

import (
	"fmt"
	"grpc_server/handlers"
	pb "grpc_server/template_service"
	"log"
	"net"

	"google.golang.org/grpc"
)

func newServer() *handlers.TemplateServiceServer {
	s := &handlers.TemplateServiceServer{}
	return s
}

// Starts the GRPC server
func main() {

	// if we are using in production, make this a flag, or read from an environmental variable
	port := 9090

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	// }
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterTemplateServiceServer(grpcServer, newServer())
	log.Println("Starting gusgrpc-service server on port:", port)

	// this is a blocking call, the program will not exit until the grpc server is killed
	// for production, want to use http endpoint to receive shutdown/restart requests
	grpcServer.Serve(lis)

}
