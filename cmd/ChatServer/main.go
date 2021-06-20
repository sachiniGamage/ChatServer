package main

import (
	"fmt"
	"log"
	"net"

	"github.com/SachiniGamage/ChatServer/internal/messaging"
	// "github.com/SachiniGamage/ChatServer/internal/messaging/cassandra"
	"github.com/SachiniGamage/ChatServer/stub"
	"google.golang.org/grpc"
)

func main() {

	// cassandraSession := cassandra.Session
	// defer cassandraSession.Close()

	//---------

	fmt.Println("hello world")
	messageService := &messaging.MessagingService{}
	registerService := &messaging.AuthenticationService{}

	port := 50052
	//listen for client request
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	fmt.Println("hello world")
	// cassandra.DBConnection()
	// cassandra.Tables()
	// cassandra.TableRegisterInsertions()
	// cassandra.Tables()

	//--------

	// // if *tls {
	// // 	if *certFile == "" {
	// // 		*certFile = data.Path("x509/server_cert.pem")
	// // 	}
	// // 	if *keyFile == "" {
	// // 		*keyFile = data.Path("x509/server_key.pem")
	// // 	}
	// // 	creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
	// // 	if err != nil {
	// // 		log.Fatalf("Failed to generate credentials %v", err)
	// // 	}
	// // 	opts = []grpc.ServerOption{grpc.Creds(creds)}
	// // }

	//-----------
	// cassandra.ChatTableInsert()

	//create an instance of the grpc server
	// grpcServer := grpc.NewServer(opts...)
	// stub.RegisterChatServiceServer(grpcServer, &messaging.MessagingService{
	// 	ChatMessage: make(map[string][]chan *stub.ChatMessageFromServer),
	// })
	// //blocking wait until process is killed
	// grpcServer.Serve(lis)

	//------------
	grpcServer := grpc.NewServer(opts...)
	stub.RegisterChatServiceServer(grpcServer, messageService)
	stub.RegisterAuthenticateUserServer(grpcServer, registerService)
	//blocking wait until process is killed
	grpcServer.Serve(lis)
	//-------------

	//cassandra

	// cluster := gocql.NewCluster("127.0.0.1")
	// cluster.Consistency = gocql.Any
	// cluster.ProtoVersion = 4
	// cluster.ConnectTimeout = time.Second * 10
	// cluster.Keyspace = "chata"
	// // cluster.Authenticator = gocql.PasswordAuthenticator{Username: "cassandra", Password: "Password"}
	// session, err := cluster.CreateSession()
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// defer session.Close()

	// err = session.Query("CREATE KEYSPACE chat WITH REPLICATION = {'class' : 'NetworkTopologyStrategy', 'datacenter1' : 3};").Exec()
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	// err = session.Query("USE chat;").Exec()
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	// err = session.Query("CREATE TABLE register( registerID int, password text, userEmail text, userName text, PRIMARY KEY(registerID));").Exec()
	// if err != nil {
	// 	log.Println(err)
	// }

	// err = session.Query("INSERT INTO register(registerid,password,useremail,username) VALUES(1,'sachi1234','sach@gmail.com','sach');").Exec()
	// err = session.Query("INSERT INTO register(registerid,password,useremail,username) VALUES(2,'sachi1234','sach@gmail.com','sach');").Exec()
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	// time_output := session.Query("SELECT * FROM register;").Iter()
	// fmt.Println("output: ", time_output)

}
