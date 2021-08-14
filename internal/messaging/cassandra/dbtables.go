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

type ChatRetrieveStruct struct {
	From    string
	To      string
	Message string
	time    int64
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
	err = session.Query("CREATE TABLE register( registerID int, password text, userEmail text,publickey text, userName text, PRIMARY KEY(userEmail,publickey));").Exec()
	err = session.Query("CREATE TABLE chatdb( registerID int, chatID uuid,msgID uuid, fromUser text, toUser text ,msg text, username text,time timestamp, PRIMARY KEY((fromUser,toUser),time)) WITH CLUSTERING ORDER BY(time ASC);").Exec()
	err = session.Query("CREATE TABLE friends(emailF1 text,myemail text, friendName text , addedEmailf1 text, addbymyemail text PRIMARY KEY(emailF1,myemail));").Exec()

	if err != nil {
		log.Println(err)
	}
	if session != nil {
		log.Println("session available")
	}
}

//insert to register
func TableRegisterInsertions(password string, email string, publickey string, username string) {
	Tables()
	if session == nil {
		log.Println("session not available")
	}
	log.Println("session available")
	err := session.Query("INSERT INTO register(registerid,password,useremail,publickey,username) VALUES(3,?,?,?,?);", password, email, publickey, username).Exec()

	if err != nil {
		log.Println(err)
		return
	}

	time_output := session.Query("SELECT * FROM register;").Iter()
	fmt.Println("output: ", time_output)
}

//insert to chat db
//([]string)
func ChatTableInsert(fromUser string, toUser string, sendmsg string) []string {
	Tables()

	var (
		msg      string
		msg2     string
		msgArray []string
		// msgArray2 []string
		count int
	)
	err := session.Query("INSERT INTO chatdb(fromUser,toUser,time,chatid,msg,msgid,registerid,username) VALUES(?,?,toUnixTimestamp(now()), now(),?,now(), 1, 'a');", fromUser, toUser, sendmsg).Exec()
	// err := session.Query("INSERT INTO chatdb(registerid,chatid,msgid,fromUser,toUser,msg,username,time) VALUES(1,1,now(),?,?,?,'a',toUnixTimestamp(now()));", fromUser, toUser, sendmsg).Exec()
	iter := session.Query("Select msg from chatdb where fromuser = ? and touser = ? order by time ALLOW FILTERING;", fromUser, toUser).Iter()

	iter2 := session.Query("Select msg from chatdb where  touser = ? and fromuser = ? order by time ALLOW FILTERING;", toUser, fromUser).Iter()

	scanner := iter.Scanner()
	scanner2 := iter2.Scanner()

	for scanner.Next() {
		err := scanner.Scan(&msg)
		if err != nil {
			log.Println("iter ")
		}
		msgArray = append(msgArray, msg)

		count++
		if count == 50 {
			break
		}
	}
	// iter.Scan(&msg)
	// iter2.Scan(&msg2)
	count = 0
	for scanner2.Next() {
		err := scanner2.Scan(&msg2)
		if err != nil {
			log.Println("iter ")
		}
		msgArray = append(msgArray, msg2)
		count++
		if count == 50 {
			break
		}
		log.Println(msgArray)
		// if iter2 != nil {
		// 	log.Println("iter 2 : " + msg2)
		// 	return msg2
		// }

	}
	if err != nil {
		log.Println(err)
		// return
	}

	return msgArray
}

func ChatRetrieve(user string) []ChatRetrieveStruct {
	Tables()

	var (
		msg       string
		fromuser  string
		touser    string
		time      int64
		msg2      string
		msgArray1 []ChatRetrieveStruct
		msgArray2 []ChatRetrieveStruct
		count     int
	)

	iter := session.Query("Select msg,fromuser,touser,time from chatdb where fromuser = ? ALLOW FILTERING;", user).Iter()

	iter2 := session.Query("Select msg,fromuser,touser,time from chatdb where  touser = ?  ALLOW FILTERING;", user).Iter()

	scanner := iter.Scanner()
	scanner2 := iter2.Scanner()

	for scanner.Next() {
		err := scanner.Scan(&msg, &fromuser, &touser, &time)
		if err != nil {
			log.Println("iter ")

		}
		log.Println("timeee")
		log.Println(&time)
		msgArray1 = append(msgArray1, ChatRetrieveStruct{From: fromuser, To: touser, Message: msg, time: time})
		log.Println("msgArray1:")
		log.Println(msgArray1)

		count++
		if count == 50 {
			break
		}
	}
	// iter.Scan(&msg)
	// iter2.Scan(&msg2)
	count = 0
	for scanner2.Next() {
		err := scanner2.Scan(&msg2, &fromuser, &touser, &time)
		if err != nil {
			log.Println("iter ")
		}
		msgArray2 = append(msgArray2, ChatRetrieveStruct{From: fromuser, To: touser, Message: msg2, time: time})
		log.Println("msgArray2:  ")

		log.Println(msgArray2)

		count++
		if count == 50 {
			break
		}
		// if iter2 != nil {
		// 	log.Println("iter 2 : " + msg2)
		// 	return msg2
		// }

		// var recMsg ChatRetrieveStruct
		// recMsg = ChatRetrieveStruct{
		// 	from: msg,
		// 	to:   msg2,
		// }
	}

	//merge sort

	var (
		ArrLen1 int
		ArrLen2 int
		i       int
		j       int
	)

	ArrLen1 = len(msgArray1)
	ArrLen2 = len(msgArray2)
	i = 0
	j = 0

	var newArr []ChatRetrieveStruct

	for i < ArrLen1 && j < ArrLen2 {
		if msgArray1[i].time >= msgArray2[j].time {
			newArr = append(newArr, msgArray2[j])
			j += 1

		} else {
			newArr = append(newArr, msgArray1[i])
			i += 1
		}
		log.Println("newArr ")
		log.Println(newArr)
	}
	for i < ArrLen1 {
		newArr = append(newArr, msgArray1[i])
		i += 1
	}

	for j < ArrLen2 {
		newArr = append(newArr, msgArray2[j])
		j += 1
	}
	return newArr
}

func GetMsg() {
	Tables()
	if session == nil {
		log.Println("session not available")
	}
	log.Println("session available")

	// iter := session.Query("select msg from chatdb where fromuser='s' and touser='u' ;").Iter()
	// iter1:= session.Query("select msg from chatdb where fromuser='u' and touser='s' ;").Iter()

}

//login
func Login(email string, passwrd string) string {
	log.Println("login3")
	var publickey string
	Tables()
	if session == nil {
		log.Println("session not available")
	}
	log.Println("session available")

	iter := session.Query("SELECT useremail,password FROM register WHERE useremail = ? AND password = ? ALLOW FILTERING;", email, passwrd).Iter()
	// err = session.Query("SELECT * FROM register where password = '?' ;", password).Exec()
	log.Println("login4")
	iter2 := session.Query("SELECT publickey FROM register where useremail = ? ALLOW FILTERING;", email).Iter()

	if (iter.NumRows()) == 0 {
		log.Println("user not found")
		return ""
	} else if (iter.NumRows()) == 1 {
		log.Println("user found")
		return "true"

	}

	if (iter2.NumRows()) == 0 {
		log.Println("No such friend")
		return ""
	} else if (iter2.NumRows()) == 1 {
		iter2.Scan(&publickey)
	}

	if iter != nil && iter2 != nil {

		log.Println(iter)
		return "true"
	}
	time_output := session.Query("SELECT * FROM register;").Iter()
	fmt.Println("output: ", time_output)

	log.Println("publickey " + publickey)
	return publickey
}

func UpdateName(name string) {
	Tables()

	if session == nil {
		log.Println("session not available")
	}
	log.Println("session available")
	err := session.Query("UPDATE register SET username = ? where useremail = 's' ;", name).Exec()
	log.Println("Name updated")
	if err != nil {
		log.Println(err)
		return
	}

	time_output := session.Query("SELECT * FROM register;").Iter()
	fmt.Println("output: ", time_output)

}

func AddFriend(emailf string, myemail string, addedEmailf1 string, addbymyemail string) [2]string {
	Tables()
	var (
		username  string
		publickey string
		newArray  [2]string
		emtyArr   [2]string
	)
	if session == nil {
		log.Println("session not available")
	}
	log.Println("session available")

	iter := session.Query("SELECT username FROM register where useremail = ? ALLOW FILTERING;", emailf).Iter()
	iter2 := session.Query("SELECT publickey FROM register where useremail = ? ALLOW FILTERING;", emailf).Iter()
	if (iter.NumRows()) == 0 {
		log.Println("No such friend")
		return emtyArr
	} else if (iter.NumRows()) == 1 {

		iter.Scan(&username)
		newArray[0] = username
		err := session.Query("INSERT INTO friends(emailF1,myemail,addedEmailf1,addbymyemail,friendName) VALUES(?,?,?,?,?);", emailf, myemail, addedEmailf1, addbymyemail, username).Exec()
		log.Println("Friend added")
		iter.Scan()
		if err != nil {
			log.Println(err)

			return newArray
		}
		// return newArray
	}

	if (iter2.NumRows()) == 0 {
		log.Println("No such friend")
		return emtyArr
	} else if (iter.NumRows()) == 1 {
		iter2.Scan(&publickey)
		newArray[1] = publickey

		iter2.Scan()

		// return newArray
	}

	if iter != nil && iter2 != nil {

		log.Println(iter)
		return newArray
	}

	if iter2 != nil {

		log.Println(iter)
		return newArray
	}
	time_output := session.Query("SELECT * FROM register;").Iter()
	fmt.Println("output: ", time_output)

	fmt.Println("newArray: ", newArray)
	return newArray
}

func ViewFriendList(email string) [][2]string {
	Tables()

	var friends [][2]string

	if session == nil {
		log.Println("session not available")
	}
	log.Println("session available")

	Scanner := session.Query("SELECT emailF1,friendname FROM friends where myemail= ? ALLOW FILTERING;", email).Iter().Scanner()

	for Scanner.Next() {

		var (
			friendEntry [2]string
		)
		Scanner.Scan(&friendEntry[0], &friendEntry[1])

		friends = append(friends, friendEntry)

	}

	time_output := session.Query("SELECT * FROM friends;").Iter()
	fmt.Println("output: ", time_output)
	return friends
}
