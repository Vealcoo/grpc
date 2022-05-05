package main

import (
	"context"
	"fmt"
	"log"
	"net"

	proto "grpc/proto"

	"google.golang.org/grpc"
)

type Server struct{}

func (*Server) Sum(ctx context.Context, req *proto.SumRequest) (*proto.SumResponse, error) {
	fmt.Printf("Sum function is invoked with %v \n", req)

	a := req.GetA()
	b := req.GetB()

	res := &proto.SumResponse{
		Result: a + b,
	}

	return res, nil
}

func main() {
	fmt.Println("starting gRPC server...")

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v \n", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterSumServiceServer(grpcServer, &Server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v \n", err)
	}
}
