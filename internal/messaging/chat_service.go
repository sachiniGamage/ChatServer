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

var values []string
var doOnce map[string]sync.Once
var channelMap map[string]chan RecieveMsg

func init() {
	channelMap = make(map[string]chan RecieveMsg)
}

func chatRecieve(sendStream stub.ChatService_ChatServer, user string) {
	// currentOnce := doOnce[user]
	log.Println("start a chat receiveing channel for user: " + user)
	// currentOnce.Do(func() {
	c, ok := channelMap[user]
	if !ok {
		c = make(chan RecieveMsg)
		channelMap[user] = c
	}
	userRelatedMsgs := cassandra.ChatRetrieve(user)

	for _, s := range userRelatedMsgs {
		// if s.To != user {
		// 	continue
		// }

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
			// From: &stub.ChatMessage{},
		}

		msg.Message.Message = i.message
		log.Println("Received msg: " + msg.Message.Message + "from :" + i.from + " is forwarded to the the user: " + user)
		if sendErr := sendStream.Send(&msg); sendErr != nil {
			fmt.Println(sendErr)
		}
	}
	// })

}

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
		// go chatRecieve(stream, in.To)
		toChannel, ok := channelMap[in.To]
		if !ok {
			toChannel = make(chan RecieveMsg, 1000)
			fmt.Println(" To Channel is created for : " + in.To)
			channelMap[in.To] = toChannel

		}
		// fmt.Println("not ok")

		// if in == nil {
		// 	continue
		// }
		// fmt.Println("Inbound Message is " + in.From + ":" + in.Message + " To: " + in.To)
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
		// msg.Message.To = in.To
		msg.Message.Message = in.Message

		if sendErr := stream.Send(&msg); sendErr != nil {
			return sendErr
		}
	}
}

func (s *AuthenticationService) Register(ctx context.Context, in *stub.RegisterUser) (*emptypb.Empty, error) {

	fmt.Println("Register Function Triggered.")
	reg := stub.RegisterUser{}
	reg.Email = "received email : " + in.Email

	cassandra.TableRegisterInsertions(in.Password, in.Email, in.Username)

	return new(emptypb.Empty), nil

}

func (s *AuthenticationService) Login(ctx context.Context, in *stub.LoginUser) (*stub.Token, error) {
	fmt.Println("Login Function Triggered.")
	login := stub.LoginUser{}
	login.Email = "email : " + in.Email

	// cassandra.Login(in.Email, in.Password)

	login.Password = "password checked"

	// jwtCreds, err := oauth.NewServiceAccountFromFile(*serviceAccountKeyFile, *oauthScope)
	// if err != nil {
	// log.Fatalf("Failed to create JWT credentials: %v", err)
	// }
	// conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")), grpc.WithPerRPCCredentials(jwtCreds))

	if cassandra.Login(in.Email, in.Password) == true {
		return &stub.Token{
			Token: "abc",
		}, nil
	} else {
		return &stub.Token{
			Token: "",
		}, nil
	}
}

func (s *EditService) UpdateName(ctx context.Context, in *stub.Edit) (*stub.RegisterUser, error) {
	fmt.Println("Update name Function Triggered.")

	updt := stub.Edit{}
	updt.Username = "received name : " + in.Username
	cassandra.UpdateName(in.Username)

	return &stub.RegisterUser{}, nil
}

func (s *EditService) AddFriend(ctx context.Context, in *stub.AddFriendReq) (*stub.AddFriendReq, error) {
	fmt.Println("Add friend Function Triggered.")

	updt := stub.AddFriendReq{
		Detail: &stub.FriendList{
			Username: in.Detail.Username,
		},
	}

	// cassandra.AddFriend(in.FriendsEmail)

	// var getname string

	getname := cassandra.AddFriend(in.Detail.FriendsEmail, in.Myemail)

	if getname != "" {
		regUser := &stub.RegisterUser{
			Username: getname,
		}
		updt.Detail.Username = regUser
		log.Println("friend added - chatservice.go")
		log.Println(updt.Detail.Username)
		return &updt, nil
	} else {
		return nil, nil
	}

}

func (s *EditService) GetFriends(ctx context.Context, in *stub.ViewFriends) (*stub.ViewFriends, error) {
	fmt.Println("get friends Function Triggered.")

	getFrnd := stub.ViewFriends{
		// friendsInList : &stub.RegisterUser{},
	}
	getFrnd.FriendsInList = in.FriendsInList
	frientArray := cassandra.ViewFriendList(in.Myemail)

	var regUserArray []*stub.RegisterUser
	for _, friendEntry := range frientArray {
		regUser := &stub.RegisterUser{}
		regUser.Email = friendEntry[0]
		regUser.Username = friendEntry[1]
		regUserArray = append(regUserArray, regUser)
	}
	finalFriendList := &stub.ViewFriends{
		FriendsInList: regUserArray,
	}
	return finalFriendList, nil
}
