package main

import (
	"context"
	"fmt"
	"messenger/grpc/server/pkg/protocol/grpc"
	"messenger/grpc/server/pkg/service/v1"
	"os"
)

func main() {
	if err := grpc.RunServer(context.Background(), v1.NewChatServiceServer(), "3000"); err != nil {
		_, err2 := fmt.Fprintf(os.Stderr, "%v\n\n", err)
		if err2 != nil {
			return
		}
		os.Exit(1)
	}
}
