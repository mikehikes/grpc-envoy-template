package handlers

import (
	"context"
	pb "grpc_server/template_service"
)

func HelloWorld(ctx context.Context, request *pb.HelloWorldRequest) (*pb.HelloWorldResponse, error) {

	return &pb.HelloWorldResponse{
		Message: "Hello " + request.Name,
	}, nil

}
