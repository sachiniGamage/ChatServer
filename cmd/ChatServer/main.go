package main

import (
	"fmt"
	"log"
	"net"

	"github.com/SachiniGamage/ChatServer/internal/messaging"
	"github.com/SachiniGamage/ChatServer/stub"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("hello world")
	messageService := &messaging.MessagingService{}

	port := 50051

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
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
	stub.RegisterChatServiceServer(grpcServer, messageService)
	grpcServer.Serve(lis)

}
