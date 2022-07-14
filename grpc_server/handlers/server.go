package handlers

import (
	pb "grpc_server/template_service"
)

type TemplateServiceServer struct {
	pb.UnimplementedTemplateServiceServer
}
