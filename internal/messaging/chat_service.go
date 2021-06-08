package messaging

import (
	"fmt"
	"io"

	"github.com/SachiniGamage/ChatServer/stub"
)

type MessagingService struct {
	stub.UnimplementedChatServiceServer
	id string
}

func (s *MessagingService) Chat(stream stub.ChatService_ChatServer) error {
	fmt.Println("Function Triggered.")
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		fmt.Println("Inbound Message is " + in.From + ":" + in.Message)
		msg := stub.ChatMessageFromServer{
			Message: &stub.ChatMessage{},
		}
		msg.Message.Message = "received " + in.Message
		if sendErr := stream.Send(&msg); sendErr != nil {
			return sendErr
		}
	}
}
