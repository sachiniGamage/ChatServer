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
	registerService := &messaging.AuthenticationService{}
	updateService := &messaging.EditService{}

	port := 50052
	//listen for client request
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	fmt.Println("hello world")

	grpcServer := grpc.NewServer(opts...)
	stub.RegisterChatServiceServer(grpcServer, messageService)
	stub.RegisterAuthenticateUserServer(grpcServer, registerService)
	stub.RegisterUpdateUserServer(grpcServer, updateService)
	//blocking wait until process is killed
	grpcServer.Serve(lis)

}
