// package main

// import (
// 	"log"
// 	"net"

// 	pb "github.com/lsnascimento/grpc-go-example/greeter"
// 	"google.golang.org/grpc"
// )

// type server struct {
// 	pb.UnimplementedGreeterServer
// }

// func main() {
// 	lis, err := net.Listen("tcp", ":50051")
// 	if err != nil {
// 		log.Fatalf("falha ao ouvir: %v", err)
// 	}
// 	s := grpc.NewServer()
// 	pb.RegisterGreeterServer(s, &server{})
// 	log.Printf("Servidor gRPC ouvindo na porta :50051")
// 	s.Serve(lis)
// }
