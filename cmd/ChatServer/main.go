package main

import (
	"fmt"

	"github.com/SachiniGamage/ChatServer/internal/messaging/cassandra"
)

func main() {
	fmt.Println("hello world")
	cassandra.DBConnection()
	cassandra.Tables()
	cassandra.TableInsertions()

	// fmt.Println("hello world")
	// messageService := &messaging.MessagingService{}

	// port := 50052

	// lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	// if err != nil {
	// 	log.Fatalf("failed to listen: %v", err)
	// }
	// var opts []grpc.ServerOption
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
	// grpcServer := grpc.NewServer(opts...)
	// stub.RegisterChatServiceServer(grpcServer, messageService)
	// grpcServer.Serve(lis)

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
