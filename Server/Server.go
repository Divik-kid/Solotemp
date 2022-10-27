package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"sync"

	gRPC "github.com/Divik-kid/Solotemp/tree/mainproto"

	"google.golang.org/grpc"
)
type Server struct {
	// an interface that the server needs to have
	gRPC.UnimplementedTemplateServer

	// here you can impliment other fields that you want
}

func main() {
	//start listening on port 6000
	lis, err := net.Listen("tcp", ":6000")
	if (error != nil){
		log.Fatalf("Failed to listen on port 6000: %v", error)
	}

	grpcServer := grpc.NewServer()

	if (error != nil){
		log.Fatalf("Failed to listen on port 6000: %v", error)
	}
}

func (s *Server) SendPost(ctx context.Context, Message *message) (*Message, error) {
    //some code here
    ...
    ack :=  // make an instance of your return type
    return (ack, nil)
}