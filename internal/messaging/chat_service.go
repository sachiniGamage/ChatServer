package messaging

import (
	"context"
	"fmt"

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

		cassandra.ChatTableInsert(in.Message)
		// cassandra.TableRegisterInsertions(in.From)

		msg.Message.Message = "received " + in.Message
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

func (s *EditService) AddFriend(ctx context.Context, in *stub.FriendList) (*stub.FriendList, error) {
	fmt.Println("Add friend Function Triggered.")

	updt := stub.FriendList{
		Username: &stub.RegisterUser{},
	}
	updt.Username = in.Username
	cassandra.AddFriend(in.FriendsEmail)

	return &updt, nil
}
