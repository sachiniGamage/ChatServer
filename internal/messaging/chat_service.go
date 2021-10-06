package messaging

import (
	"context"
	"fmt"
	"log"
	"sync"

	// "go/token"
	"io"

	// "log"

	"github.com/SachiniGamage/ChatServer/internal/messaging/cassandra"
	"github.com/SachiniGamage/ChatServer/stub"

	// "google.golang.org/grpc"
	// "google.golang.org/grpc/credentials/oauth"

	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type MessagingService struct {
	stub.UnimplementedChatServiceServer
	id          string
	ChatMessage map[string][]chan *stub.ChatMessageFromServer
}

type AuthenticationService struct {
	stub.UnimplementedAuthenticateUserServer
}

type jwt struct {
	token string
}

type EditService struct {
	stub.UnimplementedUpdateUserServer
}

type RecieveMsg struct {
	from    string
	message string
	to      string
}

type RecieveGrpMsg struct {
	friendEmail string
	Msg         string
	GroupID     string
	Time        int64
}

var values []string
var doOnce map[string]sync.Once
var channelMap map[string]chan RecieveMsg
var channelMap2 map[string]chan RecieveGrpMsg

func init() {
	channelMap = make(map[string]chan RecieveMsg)
	channelMap2 = make(map[string]chan RecieveGrpMsg)
}

//get group chat - from user
func GroupChatRetrieve(sendStream stub.ChatService_GroupChatServer, user string) {

	log.Println("start a group chat receiveing channel for groupID: " + user)
	c, ok := channelMap2[user]
	if !ok {
		c = make(chan RecieveGrpMsg)
		channelMap2[user] = c
	}
	groupIDRelatedMsgs := cassandra.GroupChatRetrieve(user)

	//only return one argument by setting range
	for _, s := range groupIDRelatedMsgs {
		msg := stub.GroupMessageFromServer{
			GroupList: &stub.GroupMessage{
				GroupDetails: &stub.MakeGroup{
					GroupId:     s.GroupID,
					FriendEmail: s.Friendemail,
				},
				Msg: s.Message,
			},
		}

		if sendErr := sendStream.Send(&msg); sendErr != nil {
			fmt.Println(sendErr)
		}
	}

	for i := range c {
		msg := stub.GroupMessageFromServer{
			GroupList: &stub.GroupMessage{
				GroupDetails: &stub.MakeGroup{
					GroupId:     i.GroupID,
					FriendEmail: i.friendEmail,
				},
				Msg: i.Msg,
			},
		}
		msg.GroupList.GroupDetails.FriendEmail = i.friendEmail
		msg.GroupList.GroupDetails.GroupId = i.GroupID
		msg.GroupList.Msg = i.Msg

		log.Println("Received msg - grp: " + msg.GroupList.Msg + "from :" + msg.GroupList.GroupDetails.FriendEmail + " is forwarded to the the user: " + user)
		if sendErr := sendStream.Send(&msg); sendErr != nil {
			fmt.Println(sendErr)
		}

	}

}

//get private chat messages
func chatRecieve(sendStream stub.ChatService_ChatServer, user string) {

	log.Println("start a chat receiveing channel for user: " + user)

	c, ok := channelMap[user]
	if !ok {
		c = make(chan RecieveMsg)
		channelMap[user] = c
	}
	userRelatedMsgs := cassandra.ChatRetrieve(user)

	for _, s := range userRelatedMsgs {

		msg := stub.ChatMessageFromServer{
			Message: &stub.ChatMessage{
				Message: s.Message,
				From:    s.From,
				To:      s.To,
			},
		}
		msg.Message.Message = s.Message
		log.Println("ChatReceived msg: " + msg.Message.Message + " from :" + s.From + " is forwarded to the the user: " + user)
		if sendErr := sendStream.Send(&msg); sendErr != nil {
			fmt.Println(sendErr)
		}

	}

	for i := range c {

		msg := stub.ChatMessageFromServer{
			Message: &stub.ChatMessage{
				Message: i.message,
				From:    i.from,
			},
		}

		msg.Message.Message = i.message
		log.Println("Received msg: " + msg.Message.Message + "from :" + i.from + " is forwarded to the the user: " + user)
		if sendErr := sendStream.Send(&msg); sendErr != nil {
			fmt.Println(sendErr)
		}
	}
	// })

}

//group chat
func (s *MessagingService) GroupChat(stream stub.ChatService_GroupChatServer) error {
	md, ok := metadata.FromIncomingContext(stream.Context())
	if ok {
		values = md.Get("fromuser")
		fmt.Println(values)
	}
	log.Println("Chat method invocation from: " + values[0])

	go GroupChatRetrieve(stream, values[0])

	fmt.Println("GroupChat Function Triggered.")
	for {

		in, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("stream has ended")
			return nil
		}
		if err != nil {
			fmt.Println(" Error received from channel from user stream: ")
			return err
		}
		fmt.Println("Group Chat message is recieved: " + in.Msg + " From: " + in.FriendEmail)

		userArr := cassandra.GroupUsers(in.GroupDetails.GroupId)
		for i := 0; i < len(userArr); i++ {
			if userArr[i] == in.GroupDetails.FriendEmail {
				continue
			}
			GroupIDhannel, ok := channelMap2[userArr[i]]
			if !ok {
				//make a channel
				GroupIDhannel = make(chan RecieveGrpMsg, 1000)
				fmt.Println(" To Channel is created for : " + userArr[i])
				channelMap2[userArr[i]] = GroupIDhannel

			}

			var recMsg RecieveGrpMsg
			recMsg = RecieveGrpMsg{
				friendEmail: in.GroupDetails.FriendEmail,
				Msg:         in.Msg,
				GroupID:     in.GroupDetails.GroupId,
			}
			GroupIDhannel <- recMsg
		}

		msg := stub.GroupMessageFromServer{
			GroupList: &stub.GroupMessage{
				GroupDetails: &stub.MakeGroup{},
			},
		}

		log.Println("Add message: " + in.Msg + " From: " + in.GroupDetails.FriendEmail)

		//Todo: Add in.randomuuid to first parameter
		cassandra.GroupChatTableInsert(in.GroupDetails.FriendEmail, in.GroupDetails.GroupId, in.Msg)

		log.Println("cht")

		msg.GroupList.Msg = in.Msg
		msg.GroupList.GroupDetails.FriendEmail = in.GroupDetails.FriendEmail
		msg.GroupList.GroupDetails.GroupId = in.GroupDetails.GroupId
		log.Println(msg.GroupList.Msg)
		log.Println(msg.GroupList.GroupDetails.FriendEmail)
		log.Println(msg.GroupList.GroupDetails.GroupId)
		log.Println(msg)
		if sendErr := stream.Send(&msg); sendErr != nil {
			return sendErr
		}
	}
}

//private chat
func (s *MessagingService) Chat(stream stub.ChatService_ChatServer) error {

	md, ok := metadata.FromIncomingContext(stream.Context())
	if ok {
		values = md.Get("fromuser")
		fmt.Println(values)
	}

	log.Println("Chat method invocation from: " + values[0])

	go chatRecieve(stream, values[0])

	for {

		in, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("stream has ended")
			return nil
		}
		if err != nil {
			fmt.Println(" Error received from channel from user stream: "+values[0], err)
			return err
		}
		fmt.Println("Chat message is recieved from chat method: " + in.Message + " From: " + in.From + " To: " + in.To + " From User stream: " + values[0])

		toChannel, ok := channelMap[in.To]
		if !ok {
			toChannel = make(chan RecieveMsg, 1000)
			fmt.Println(" To Channel is created for : " + in.To)
			channelMap[in.To] = toChannel

		}
		msg := stub.ChatMessageFromServer{
			Message: &stub.ChatMessage{},
		}

		var recMsg RecieveMsg
		recMsg = RecieveMsg{
			from:    in.From,
			message: in.Message,
			to:      in.To,
		}
		toChannel <- recMsg

		log.Println("Add message: " + in.Message + " From: " + in.From + " to receiveMessage Channel")

		cassandra.ChatTableInsert(in.From, in.To, in.Message)

		msg.Message.From = in.From
		msg.Message.Message = in.Message

		if sendErr := stream.Send(&msg); sendErr != nil {
			return sendErr
		}
	}
}

//register user
func (s *AuthenticationService) Register(ctx context.Context, in *stub.RegisterUser) (*emptypb.Empty, error) {

	fmt.Println("Register Function Triggered.")
	reg := stub.RegisterUser{}
	reg.Email = "received email : " + in.Email

	cassandra.TableRegisterInsertions(in.Password, in.Email, in.PublicKey, in.Username)

	return new(emptypb.Empty), nil

}

//login user
func (s *AuthenticationService) Login(ctx context.Context, in *stub.LoginUser) (*stub.Token, error) {
	fmt.Println("Login Function Triggered.")
	login := stub.LoginUser{}
	login.Email = "email : " + in.Email

	login.Password = "password checked"

	pbkey := cassandra.Login(in.Email, in.Password)
	fmt.Println("pbkey:" + pbkey)
	if pbkey != "" {
		return &stub.Token{
			Token:     "abc",
			PublicKey: pbkey,
		}, nil
	} else {
		return &stub.Token{
			Token: "",
		}, nil
	}
}

//update profile user name
func (s *EditService) UpdateName(ctx context.Context, in *stub.Edit) (*stub.RegisterUser, error) {
	fmt.Println("Update name Function Triggered.")

	updt := stub.Edit{}
	updt.Username = "received name : " + in.Username
	//parameter 2t myemail
	cassandra.UpdateName(in.Username, in.Username)

	return &stub.RegisterUser{}, nil
}

func (s *EditService) CreateGroup(ctx context.Context, in *stub.MakeGroup) (*stub.MakeGroup, error) {
	fmt.Println("Create Group Function Triggered.")
	grp := stub.MakeGroup{
		GroupId:     in.GroupId,
		GroupName:   in.GroupName,
		AdminEmail:  in.AdminEmail,
		FriendEmail: in.FriendEmail,
	}
	grp.GroupName = "Group Name: " + in.GroupName
	getGroupDetails := cassandra.GroupChatDetailsInsertion(in.GroupId, in.GroupName, in.AdminEmail, in.FriendEmail)

	var emptyArr [4]string
	if getGroupDetails != emptyArr {
		details := &stub.MakeGroup{
			GroupId:     getGroupDetails[0],
			GroupName:   getGroupDetails[1],
			AdminEmail:  getGroupDetails[2],
			FriendEmail: getGroupDetails[3],
		}
		grp.GroupId = details.GroupId
		grp.GroupName = details.GroupName
		grp.AdminEmail = details.AdminEmail
		grp.FriendEmail = details.FriendEmail
		return &grp, nil
	} else {
		return nil, nil
	}

}

//Add a friend - private chat
func (s *EditService) AddFriend(ctx context.Context, in *stub.AddFriendReq) (*stub.AddFriendReq, error) {
	fmt.Println("Add friend Function Triggered.")

	updt := stub.AddFriendReq{
		Addbymyemail: in.Addbymyemail,

		Detail: &stub.FriendList{
			Username:  in.Detail.Username,
			PublicKey: in.Detail.PublicKey,
		},
	}
	fmt.Println("Addbymyemail" + in.Addbymyemail)

	getname := cassandra.AddFriend(in.Detail.FriendsEmail, in.Myemail, in.AddedEmailf1, in.Addbymyemail)
	var emtyArr [3]string
	if getname != emtyArr {
		regUser := &stub.RegisterUser{
			Username:  getname[0],
			PublicKey: getname[1],
		}
		frndRq := &stub.AddFriendReq{
			Addbymyemail: getname[2],
		}
		updt.Detail.Username = regUser
		updt.Detail.PublicKey = regUser.PublicKey
		updt.Addbymyemail = frndRq.Addbymyemail
		log.Println("friend added - chatservice.go")
		log.Println(updt.Detail.Username)
		log.Println(updt.Detail.PublicKey)
		log.Println(getname[2])
		log.Println(frndRq.Addbymyemail)
		log.Println(updt.Addbymyemail)
		return &updt, nil
	} else {
		return nil, nil
	}

}

//get groups
func (s *EditService) GetGroup(ctx context.Context, in *wrapperspb.StringValue) (*stub.ViewGroup, error) {
	fmt.Println("get Groups Function Triggered.")

	grpArray := cassandra.ViewGroupList(in.Value)
	var groupArray []*stub.MakeGroup
	for _, groupEntry := range grpArray {
		grp := &stub.MakeGroup{}
		grp.GroupId = groupEntry[0]
		grp.GroupName = groupEntry[1]

		groupArray = append(groupArray, grp)
	}
	finalGroupList := &stub.ViewGroup{
		GrpDetails: groupArray,
	}
	return finalGroupList, nil
}

//get friends
func (s *EditService) GetFriends(ctx context.Context, in *stub.ViewFriends) (*stub.ViewFriends, error) {
	fmt.Println("get friends Function Triggered.")

	getFrnd := stub.ViewFriends{}
	getFrnd.FriendsInList = in.FriendsInList
	frientArray := cassandra.ViewFriendList(in.Myemail)

	var regUserArray []*stub.RegisterUser
	for _, friendEntry := range frientArray {

		regUser := &stub.RegisterUser{}
		regUser.Email = friendEntry[0]
		regUser.Username = friendEntry[1]
		regUser.EncryptedKey = friendEntry[2]

		regUserArray = append(regUserArray, regUser)
	}
	finalFriendList := &stub.ViewFriends{
		FriendsInList: regUserArray,
	}

	return finalFriendList, nil
}
