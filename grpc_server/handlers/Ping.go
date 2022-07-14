package handlers

import (
	"context"
	pb "grpc_server/template_service"
)

func Ping(ctx context.Context, request *pb.PingRequest) (*pb.PingResponse, error) {

	return &pb.PingResponse{}, nil

}
