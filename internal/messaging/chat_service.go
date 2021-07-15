package messaging

import (
	"context"
	"fmt"
	"log"

	// "go/token"
	"io"

	// "log"

	"github.com/SachiniGamage/ChatServer/internal/messaging/cassandra"
	"github.com/SachiniGamage/ChatServer/stub"

	// "google.golang.org/grpc"
	// "google.golang.org/grpc/credentials/oauth"
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

func (s *MessagingService) Chat(stream stub.ChatService_ChatServer) error {
	fmt.Println("Function Triggered.")
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("stream has ended")
			return nil
		}
		fmt.Println("msg recieved")
		if err != nil {
			fmt.Println(err)
		}
		// if in == nil {
		// 	continue
		// }
		fmt.Println("Inbound Message is " + in.From + ":" + in.Message)
		msg := stub.ChatMessageFromServer{
			Message: &stub.ChatMessage{},
		}
		// TODO: set fromuser and touser in parameters
		cassandra.ChatTableInsert(in.From, in.To, in.Message)

		msg.Message.Message = "received " + in.Message + in.From + in.To
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
