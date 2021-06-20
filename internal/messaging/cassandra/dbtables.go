package cassandra

import (
	"fmt"
	"log"
	"time"

	"github.com/gocql/gocql"
	// "google.golang.org/genproto/googleapis/bigtable/admin/cluster/v1"
	//"github.com/SachiniGamage/ChatServer/internal/messaging/cassandra"
)

type MessageManager struct {
	Session *gocql.Session
}

var cluster *gocql.ClusterConfig
var session *gocql.Session
var Scanner gocql.Scanner

func Tables() {

	cluster = gocql.NewCluster("127.0.0.1")
	cluster.Consistency = gocql.Any
	cluster.ProtoVersion = 4
	cluster.ConnectTimeout = time.Second * 10
	cluster.Keyspace = "sender"
	var err error
	session, err = cluster.CreateSession()
	if err != nil {
		log.Println(err)
		return
	}

	// err = session.Query("CREATE KEYSPACE sender WITH REPLICATION = {'class' : 'NetworkTopologyStrategy', 'datacenter1': 1};").Exec()
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	err = session.Query("CREATE TABLE register( registerID int, password text, userEmail text, userName text, PRIMARY KEY(userEmail));").Exec()
	err = session.Query("CREATE TABLE chatdb( registerID int, chatID id,msgID uuid, sendmsg text, username text,time timestamp, PRIMARY KEY(msgID));").Exec()

	if err != nil {
		log.Println(err)
	}
	if session != nil {
		log.Println("session available")
	}
}

func TableRegisterInsertions(password string, email string, username string) {
	Tables()
	if session == nil {
		log.Println("session not available")
	}
	log.Println("session available")
	err := session.Query("INSERT INTO register(registerid,password,useremail,username) VALUES(3,?,?,?);", password, email, username).Exec()

	if err != nil {
		log.Println(err)
		return
	}

	time_output := session.Query("SELECT * FROM register;").Iter()
	fmt.Println("output: ", time_output)
}

func ChatTableInsert(sendmsg string) {
	Tables()

	err := session.Query("INSERT INTO chatdb(registerid,chatid,msgid,sendmsg,username,time) VALUES(1,1,now(),?,'a',toUnixTimestamp(now()));", sendmsg).Exec()

	if err != nil {
		log.Println(err)
		return
	}

	time_output := session.Query("SELECT * FROM chatdb;").Iter()
	fmt.Println("output: ", time_output)
}

func ViewSendMessage() {

	// err := session.Query("SELECT sendmsg FROM chatdb ;").Exec()
	err := session.Query("SELECT sendmsg FROM chatdb WHERE registerid=2 AND chatid=1;").Exec()

	if err != nil {
		log.Println(err)
		return
	}

	time_output := session.Query("SELECT * FROM chatdb;").Iter()
	fmt.Println("output: ", time_output)

}

func ViewRecievedMessage() {

	// err := session.Query("SELECT sendmsg FROM chatdb ;").Exec()
	Scanner = session.Query("SELECT sendmsg FROM chatdb WHERE registerid=2 AND chatid=1;").Iter().Scanner()
	var sendMsg string
	var err error
	for Scanner.Next() {
		err = Scanner.Scan(&sendMsg)
		if err != nil {
			log.Println(err)
			return
		}
	}

	time_output := session.Query("SELECT * FROM chatdb;").Iter()
	fmt.Println("output: ", time_output)

}
