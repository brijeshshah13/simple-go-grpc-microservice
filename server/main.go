package main

import (
	"context"
	"net"

	"github.com/brijeshshah13/simple-go-grpc-microservice/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func add(ctx context.Context, request *protos.Request) (*protos.Response, error) {
	a, b := request.GetA(), request.GetB()

	result := a + b

	return &protos.Response{Result: result}, nil
}

func multiply(ctx context.Context, request *protos.Request) (*protos.Response, error) {
	a, b := request.GetA(), request.GetB()

	result := a * b

	return &protos.Response{Result: result}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()

	AlgebraServiceObj := protos.AlgebraService{
		Add:      add,
		Multiply: multiply,
	}

	protos.RegisterAlgebraService(srv, &AlgebraServiceObj)
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}

}
