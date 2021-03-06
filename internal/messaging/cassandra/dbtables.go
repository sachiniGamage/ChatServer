package cassandra

import (
	"crypto/tls"
	"fmt"
	"log"
	"path/filepath"
	"time"

	"github.com/gocql/gocql"
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

type GroupChatRetriveStruct struct {
	Friendemail string
	Message     string
	GroupID     string
	Time        int64
	AdminEmail  string
}

var cluster *gocql.ClusterConfig
var session *gocql.Session
var Scanner gocql.Scanner

func Tables() {
	var _cqlshrc_host = "86ce02b2-04e7-41e4-8278-764e9b1e5ae1-asia-south1.apps.astra.datastax.com"
	var _cqlshrc_port = "29042"
	var _username = "ssgdummy1@gmail.com"
	var _password = "YzTTkb3T#CiKy2!"
	cluster = gocql.NewCluster(_cqlshrc_host)
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: _username,
		Password: _password,
	}
	cluster.Hosts = []string{_cqlshrc_host + ":" + _cqlshrc_port}
	certPath, _ := filepath.Abs("/home/sachini/Desktop/Cassandra/cert")
	keyPath, _ := filepath.Abs("/home/sachini/Desktop/Cassandra/key")
	cert, _ := tls.LoadX509KeyPair(certPath, keyPath)

	tlsConfig := &tls.Config{
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: true,
	}
	cluster.SslOpts = &gocql.SslOptions{
		Config:                 tlsConfig,
		EnableHostVerification: false,
	}
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
	err = session.Query("CREATE TABLE friends(emailF1 text,myemail text, friendName text , addedEmailf1 text, addbymyemail text ,PRIMARY KEY(emailF1,myemail));").Exec()
	err = session.Query("CREATE TABLE grpchatdb(groupID uuid,friendEmail text,msg text,msgID uuid,time timestamp, PRIMARY KEY((groupID),time)) WITH CLUSTERING ORDER BY(time ASC);").Exec()
	err = session.Query("CREATE TABLE grpdetaildb(ID UUID, groupID text,groupName text,adminEmail text, friendEmail text, PRIMARY KEY(ID));").Exec()

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

//Group details
func GroupChatDetailsInsertion(groupID string, groupname string, adminemail string, friendEmail string) [4]string {
	Tables()
	var (
		grpID      string
		grpName    string
		adminEmail string
		frndEmail  string
		msgArray   [4]string
		emptyArr   [4]string
	)
	err := session.Query("INSERT INTO grpdetaildb(ID,groupID,groupname,adminemail,friendEmail) VALUES(now(),?,?,?,?);", groupID, groupname, adminemail, friendEmail).Exec()
	//Todo: Write a new select query to get grp chats which is i'm already in
	iter := session.Query("Select groupID,groupName,adminEmail,friendemail from grpdetaildb where AdminEmail = ?  ALLOW FILTERING;", adminemail).Iter()
	scanner := iter.Scanner()
	for scanner.Next() {
		err := scanner.Scan(&grpID, &grpName, &adminEmail, &frndEmail)
		if err != nil {
			log.Println("error in iter - groupchatDetails ")
			return emptyArr
		}
		msgArray[0] = grpID
		msgArray[1] = grpName
		msgArray[2] = adminEmail
		msgArray[3] = frndEmail
	}
	if err != nil {
		log.Println(err)
	}
	log.Println("grp msgArr")
	log.Println(msgArray)
	return msgArray
}

//message in groups
func GroupChatTableInsert(friendEmail string, groupId string, message string) [3]string {
	Tables()
	var (
		frndEmail string
		msg       string
		time      int64
		msgArray  [3]string
		emptyArr  [3]string
	)
	err := session.Query("INSERT INTO grpchatdb(groupID,friendEmail,msg,msgID,time) VALUES(?,?,?,now(),toUnixTimestamp(now()));", groupId, friendEmail, message).Exec()
	iter := session.Query("Select friendEmail,msg,time from grgpchatdb ").Iter()
	scanner := iter.Scanner()
	for scanner.Next() {
		err := scanner.Scan(&frndEmail, &msg, &time)
		if err != nil {
			log.Println("iter ")
			return emptyArr
		}
		msgArray[0] = frndEmail
		msgArray[1] = msg
	}
	if err != nil {
		log.Println(err)
	}
	return msgArray
}

//private chat - messaging
func ChatTableInsert(fromUser string, toUser string, sendmsg string) []string {
	Tables()
	var (
		msg      string
		msg2     string
		msgArray []string
		count    int
	)
	err := session.Query("INSERT INTO chatdb(fromUser,toUser,time,chatid,msg,msgid,registerid,username) VALUES(?,?,toUnixTimestamp(now()), now(),?,now(), 1, 'a');", fromUser, toUser, sendmsg).Exec()
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
		//50 msgs
		count++
		if count == 50 {
			break
		}
	}
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
	}
	if err != nil {
		log.Println(err)
	}
	return msgArray
}

func CheckId(array []string, id string) bool {
	var i int
	for i = 0; i < len(array); i++ {
		if array[i] == id {
			return true
		}
	}
	return false
}

//get group chat details
func GroupChatRetrieve(email string) []GroupChatRetriveStruct {
	Tables()
	var (
		msg         string
		friendemail string
		adminemail  string
		groupid     string
		time        int64
		count       int
		arr         []string
		msgArr      []GroupChatRetriveStruct
	)
	iter2 := session.Query("select groupid from grpdetaildb where friendemail =? ALLOW FILTERING;", email).Iter()
	iter3 := session.Query("select groupid from grpdetaildb where adminemail =? ALLOW FILTERING;", email).Iter()

	scanner2 := iter2.Scanner()
	scanner3 := iter3.Scanner()

	if scanner2 != nil || scanner3 != nil {
		for scanner2.Next() {
			var id string
			err := scanner2.Scan(&id)
			if err != nil {
				log.Println("iter ")
			}
			arr = append(arr, id)
		}
		for scanner3.Next() {
			var id string
			err := scanner3.Scan(&id)
			if err != nil {
				log.Println("iter ")
			}
			if !CheckId(arr, id) {
				arr = append(arr, id)
			}
		}
		for i := 0; i < len(arr); i++ {
			iter := session.Query("select msg,friendemail,groupid,time from grpchatdb where groupid =? ", arr[i]).Iter()
			iter4 := session.Query("select friendemail,adminemail from grpdetaildb where groupid = ?", arr[i]).Iter()

			scanner := iter.Scanner()
			scanner2 := iter4.Scanner()

			log.Println(iter)
			log.Println(arr[i])
			log.Println("grp chat**")

			if scanner != nil {
				for scanner.Next() {
					err := scanner.Scan(&msg, &friendemail, &groupid, &time)
					if err != nil {
						log.Println("iter ")
					}
					log.Println("msg")
					log.Println(msg)
					msgArr = append(msgArr, GroupChatRetriveStruct{Friendemail: friendemail, Message: msg, GroupID: groupid, Time: time})
					count++
					if count == 50 {
						break
					}
				}
			}
			if scanner2 != nil {
				for scanner2.Next() {
					err := scanner2.Scan(&adminemail)
					if err != nil {
						log.Println("iter ")
					}
					log.Println("adminemail")
					log.Println(adminemail)
					msgArr = append(msgArr, GroupChatRetriveStruct{AdminEmail: adminemail})

					count++
					if count == 50 {
						break
					}
				}
			}
		}
	}
	return msgArr
}

//get group member email
func GroupUsers(groupId string) []string {
	Tables()
	iter4 := session.Query("select friendemail,adminemail from grpdetaildb where groupid = ? allow filtering", groupId).Iter()
	scanner := iter4.Scanner()
	var userArr []string
	var adminemail string
	if scanner != nil {
		for scanner.Next() {
			var (
				friendemail string
			)
			err := scanner.Scan(&friendemail, &adminemail)
			if err != nil {
				log.Println("iter ")
			}
			log.Println("msg")
			log.Println(friendemail)
			userArr = append(userArr, friendemail)
		}
		userArr = append(userArr, adminemail)
	}
	return userArr
}

//get private chat details
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
		if (iter2.NumRows()) == 0 {
			log.Println("No such friend")
			return ""
		} else if (iter2.NumRows()) == 1 {
			iter2.Scan(&publickey)
			log.Println(publickey)
		}
		return publickey
	}
	return publickey
}

//update current user profile name
func UpdateName(name string, myemail string) {
	Tables()
	if session == nil {
		log.Println("session not available")
	}
	log.Println("session available")
	err := session.Query("UPDATE register SET username = ? where useremail = ?;", name, myemail).Exec()
	log.Println("Name updated")
	if err != nil {
		log.Println(err)
		return
	}
	time_output := session.Query("SELECT * FROM register;").Iter()
	fmt.Println("output: ", time_output)
}

//add friend - private chat
func AddFriend(emailf string, myemail string, addedEmailf1 string, addbymyemail string) [3]string {
	Tables()
	var (
		username     string
		publickey    string
		addedmyemail string
		newArray     [3]string
		emtyArr      [3]string
	)
	if session == nil {
		log.Println("session not available")
	}
	log.Println("session available")

	iter := session.Query("SELECT username FROM register where useremail = ? ALLOW FILTERING;", emailf).Iter()
	iter2 := session.Query("SELECT publickey FROM register where useremail = ? ALLOW FILTERING;", emailf).Iter()
	iter3 := session.Query("SELECT addbymyemail FROM friends where emailF1 = ? AND myemail=? ALLOW FILTERING;", emailf, myemail).Iter()
	// iter3 := session.Query("SELECT addbymyemail FROM friends where emailF1 = ? ALLOW FILTERING;", emailf).Iter()
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
	}
	if (iter2.NumRows()) == 0 {
		log.Println("No such publickey")
		return emtyArr
	} else if (iter2.NumRows()) == 1 {
		iter2.Scan(&publickey)
		newArray[1] = publickey
		iter2.Scan()
	}
	if (iter3.NumRows()) == 0 {
		log.Println("No such addbymyemail")
	} else if (iter3.NumRows()) == 1 {
		iter3.Scan(&addedmyemail)
		newArray[2] = addedmyemail
		iter3.Scan()
		return newArray
	}
	if iter != nil && iter2 != nil && iter3 != nil {
		fmt.Println("newArray: ", newArray)
		log.Println(iter)
		return newArray
	}
	time_output := session.Query("SELECT * FROM register;").Iter()
	fmt.Println("output: ", time_output)
	fmt.Println("newArray: ", newArray)
	return newArray
}

//view groups
func ViewGroupList(myEmail string) [][2]string {
	Tables()
	if session == nil {
		log.Println("session not available")
	}
	log.Println("session available")

	var groups [][2]string

	Scanner1 := session.Query("select groupid,groupname from grpdetaildb where friendemail = ? ALLOW FILTERING;", myEmail).Iter().Scanner()
	Scanner2 := session.Query("select groupid,groupname from grpdetaildb where adminemail = ?  ALLOW FILTERING;", myEmail).Iter().Scanner()

	for Scanner1.Next() {
		var (
			groupEntry [2]string
		)
		Scanner1.Scan(&groupEntry[0], &groupEntry[1])
		groups = append(groups, groupEntry)
	}
	for Scanner2.Next() {
		var (
			groupEntry [2]string
		)
		Scanner2.Scan(&groupEntry[0], &groupEntry[1])
		groups = append(groups, groupEntry)
	}
	return groups
}

//view friends as a list - private chat
func ViewFriendList(email string) [][3]string {
	Tables()
	var friends [][3]string
	if session == nil {
		log.Println("session not available")
	}
	log.Println("session available")

	Scanner := session.Query("SELECT emailF1,friendname,addbymyemail FROM friends where myemail= ? ALLOW FILTERING;", email).Iter().Scanner()

	for Scanner.Next() {
		var (
			friendEntry [3]string
		)
		Scanner.Scan(&friendEntry[0], &friendEntry[1], &friendEntry[2])
		friends = append(friends, friendEntry)
	}
	time_output := session.Query("SELECT * FROM friends;").Iter()
	fmt.Println("output: ", time_output)
	return friends
}

//update friends
func AddFriendUpdate(addedEmailf1 string, addbymyemail string, myemail string, emailf string) string {
	Tables()
	if session == nil {
		log.Println("session not available")
	}
	log.Println("session available")

	var (
		newArray string
		emtyArr  string
	)
	err := session.Query("UPDATE friends SET addedemailf1 = ? , addbymyemail = ? where myemail = ? AND emailf1=?;", addedEmailf1, addbymyemail, myemail, emailf).Exec()
	iter1 := session.Query("SELECT addbymyemail FROM friends where emailF1 = ? AND myemail = ? ALLOW FILTERING;", emailf, myemail).Iter()

	if (iter1.NumRows()) == 0 {
		log.Println("No such addbymyemail")
		return emtyArr
	} else if (iter1.NumRows()) == 1 {
		iter1.Scan(&addbymyemail)
		newArray = addbymyemail
		iter1.Scan()
	}
	if err != nil {
		log.Println(err)
		log.Println("err in add friend update")
	}
	if iter1 != nil {
		log.Println(err)
		log.Println(iter1)
		log.Println(" Add friend update : ")
		log.Println(newArray)
		return newArray
	}
	log.Println("addedemailf1 and addbymyemail updated")
	if err != nil {
		log.Println(err)
	}
	time_output := session.Query("SELECT * FROM register;").Iter()
	fmt.Println("output: ", time_output)
	return newArray
}
