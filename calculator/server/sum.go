package main

import (
	"context"
	"log"

	pb "github.com/JackConroySunLife/Go-GRPC-Project/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, input *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Sum function was invoked with %v\n", input)
	return &pb.SumResponse{
		Result: input.FirstNumber + input.SecondNumber,
	}, nil
}
