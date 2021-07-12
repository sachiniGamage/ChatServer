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
	// cluster.Consistency = gocql.Any
	cluster.Consistency = gocql.One
	cluster.ProtoVersion = 4
	cluster.ConnectTimeout = time.Second * 10
	cluster.Keyspace = "sender"
	var err error
	session, err = cluster.CreateSession()
	if err != nil {
		log.Println(err)
		return
	}
	err = session.Query("CREATE TABLE register( registerID int, password text, userEmail text, userName text, PRIMARY KEY(userEmail));").Exec()
	err = session.Query("CREATE TABLE chatdb( registerID int, chatID id,msgID uuid, sendmsg text, username text,time timestamp, PRIMARY KEY(msgID));").Exec()
	err = session.Query("CREATE TABLE friends(emailF text,PRIMARY KEY(emailF));").Exec()
	if err != nil {
		log.Println(err)
	}
	if session != nil {
		log.Println("session available")
	}
}

//insert to register
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

//insert to chat db
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

//login
func Login(email string, passwrd string) bool {
	log.Println("login3")
	Tables()
	if session == nil {
		log.Println("session not available")
	}
	log.Println("session available")

	iter := session.Query("SELECT useremail,password FROM register WHERE useremail = ? AND password = ? ALLOW FILTERING;", email, passwrd).Iter()
	// err = session.Query("SELECT * FROM register where password = '?' ;", password).Exec()
	log.Println("login4")

	if (iter.NumRows()) == 0 {
		log.Println("user not found")
		return false
	} else if (iter.NumRows()) == 1 {
		log.Println("user found")
		return true

	}

	if iter != nil {

		log.Println(iter)
		return true
	}
	time_output := session.Query("SELECT * FROM register;").Iter()
	fmt.Println("output: ", time_output)

	return true
}

func UpdateName(name string) {
	Tables()

	if session == nil {
		log.Println("session not available")
	}
	log.Println("session available")
	err := session.Query("UPDATE register SET username = ? where useremail = 's';", name).Exec()
	log.Println("Name updated")
	if err != nil {
		log.Println(err)
		return
	}

	time_output := session.Query("SELECT * FROM register;").Iter()
	fmt.Println("output: ", time_output)

}

func AddFriend(email string) {
	Tables()
	if session == nil {
		log.Println("session not available")
	}
	log.Println("session available")
	iter := session.Query("SELECT username FROM register where email = ? ;", email).Iter()

	err := session.Query("INSERT INTO friends(emailF) VALUES(?);", email).Exec()

	if err != nil {
		log.Println(err)
		// return
	}
	var useremail string

	for iter.Scan(&useremail) {
		if useremail != email {
			log.Println("incorrect")
		}
	}
	if iter != nil {

		log.Println(iter)
		return
	}
	time_output := session.Query("SELECT * FROM register;").Iter()
	fmt.Println("output: ", time_output)
}

func GetRegName() {
	Tables()
	if session == nil {
		log.Println("session not available")
	}
	log.Println("session available")

	iter := session.Query("SELECT username FROM register ;").Iter()
	if iter != nil {

		log.Println(iter)
		return
	}
	time_output := session.Query("SELECT * FROM register;").Iter()
	fmt.Println("output: ", time_output)

}
