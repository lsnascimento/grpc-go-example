package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/lsnascimento/grpc-go-example/greeter"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	message := fmt.Sprintf("Hello, %s!", req.Name)
	return &pb.HelloReply{Message: message}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("falha ao ouvir: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})

	log.Printf("Servidor gRPC ouvindo na porta :50051")

	s.Serve(lis)
}
