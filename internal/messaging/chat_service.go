package messaging

import (
	"context"
	"fmt"
	"io"

	"github.com/SachiniGamage/ChatServer/internal/messaging/cassandra"
	"github.com/SachiniGamage/ChatServer/stub"
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

	// user,err:=s.Register(ctx)
	// if err != nil {
	// 	return nil, err
	// }

	// return &emptypb.Empty{}, nil

	// return nil, nil
	return new(emptypb.Empty), nil

}
