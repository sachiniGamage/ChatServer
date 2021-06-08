package cassandra

import (
	"fmt"
	"log"

	"github.com/gocql/gocql"
	// "google.golang.org/genproto/googleapis/bigtable/admin/cluster/v1"
	//"github.com/SachiniGamage/ChatServer/internal/messaging/cassandra"
)

func Tables() {

	// register table

	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "sender"
	session, err := cluster.CreateSession()
	if err != nil {
		log.Println(err)
		return
	}
	defer session.Close()

	// err = session.Query("CREATE KEYSPACE sender WITH REPLICATION = {'class' : 'NetworkTopologyStrategy', 'datacenter1': 1};").Exec()
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	err = session.Query("CREATE TABLE register( registerID int, password text, userEmail text, userName text, PRIMARY KEY(registerID));").Exec()

	// chat table
	err = session.Query("CREATE TABLE chatdb( registerID int, recievedmsg text, sendmsg text, username text,time timeuuid, PRIMARY KEY(registerID));").Exec()

	if err != nil {
		log.Println(err)
	}
}

func TableInsertions() {
	DBConnection()

	cluster := gocql.NewCluster("127.0.0.1")
	session, err := cluster.CreateSession()
	if err != nil {
		log.Println(err)
		return
	}
	defer session.Close()
	Tables()
	err = session.Query("INSERT INTO sender.register(registerid,password,useremail,username) VALUES(1,'sachi1234','sach@gmail.com','sach');").Exec()
	err = session.Query("INSERT INTO sender.register(registerid,password,useremail,username) VALUES(2,'sachi1234','sach@gmail.com','sach');").Exec()

	err = session.Query("INSERT INTO sender.chatdb(registerid,recievedmsg,sendmsg,username,time) VALUES(2,'sachi1234','sach@gmail.com','sach',50554d6e-29bb-11e5-b345-feff819cdc9f);").Exec()

	if err != nil {
		log.Println(err)
		return
	}

	time_output := session.Query("SELECT * FROM register;").Iter()
	fmt.Println("output: ", time_output)
}
