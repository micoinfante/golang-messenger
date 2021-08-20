package v1

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"log"
	"messenger/grpc/server/pkg/api/v1"
)

type chatServiceServer struct {
	message chan string
}

func (c chatServiceServer) Send(ctx context.Context, message *wrapperspb.StringValue) (*emptypb.Empty, error) {
	if message != nil {
		log.Printf("Send requested: message=%v\n", message)
		c.message <- message.Value
	} else {
		log.Println("Send requested: message=<empty>")
	}

	return &empty.Empty{}, nil
}

func (c chatServiceServer) Subscribe(empty *emptypb.Empty, server v1.ChatService_SubscribeServer) error {
	log.Print("Subscribe requested \n")
	for {
		m := <-c.message
		n := v1.Message{Text: fmt.Sprintf("I have received from you: %s. Thanks!", m)}
		if err := server.Send(&n); err != nil {
			c.message <- m
			log.Printf("Stream connection failed: %v", err)
			return nil
		}
		log.Printf("Message sent: %+v", n.Text)
	}
}

func NewChatServiceServer() v1.ChatServiceServer {
	return &chatServiceServer{message: make(chan string, 1000)}
}
