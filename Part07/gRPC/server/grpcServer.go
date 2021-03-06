package main

import (
	"fmt"
	"net"

	p "github.com/knix008/protobuf"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type MessageServer struct {
}

var port = ":8080"

func (MessageServer) SayIt(ctx context.Context, r *p.Request) (*p.Response,
	error) {
	fmt.Println("Request Text:", r.Text)
	fmt.Println("Request SubText:", r.Subtext)
	response := &p.Response{
		Text:    r.Text,
		Subtext: "Got it!",
	}
	return response, nil
}

func main() {
	server := grpc.NewServer()
	var messageServer MessageServer
	p.RegisterMessageServiceServer(server, messageServer)
	listen, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Serving requests...")
	server.Serve(listen)
}
