package main

import (
	"fmt"
	pb "grpc_server/template_service"
	"log"
	"net"

	"google.golang.org/grpc"
)

type TemplateServiceServer struct {
	pb.UnimplementedTemplateServiceServer
}

func newServer() *TemplateServiceServer {
	s := &TemplateServiceServer{}
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

	// for production, recommended to use https to encrypt comms between the envoy proxy
	// if *tls {
	// 	if *certFile == "" {
	// 		*certFile = data.Path("x509/server_cert.pem")
	// 	}
	// 	if *keyFile == "" {
	// 		*keyFile = data.Path("x509/server_key.pem")
	// 	}
	// 	creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
	// 	if err != nil {
	// 		log.Fatalf("Failed to generate credentials %v", err)
	// 	}
	// 	opts = []grpc.ServerOption{grpc.Creds(creds)}
	// }
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterTemplateServiceServer(grpcServer, newServer())
	log.Println("Starting gusgrpc-service server on port:", port)

	// this is a blocking call, the program will not exit until the grpc server is killed
	// for production, want to use http endpoint to receive shutdown/restart requests
	grpcServer.Serve(lis)

}
