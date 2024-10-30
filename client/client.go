// package main

// import (
// 	"context"
// 	"log"

// 	pb "github.com/lsnascimento/grpc-go-example/greeter"
// 	"google.golang.org/grpc"
// )

// func main() {
// 	conn, _ := grpc.Dial("localhost:50051", grpc.WithInsecure())
// 	client := pb.NewGreeterClient(conn)
// 	response, _ := client.SayHello(context.Background(), &pb.HelloRequest{Name: "World"})
// 	log.Printf("Resposta: %s", response.Message)
// }
