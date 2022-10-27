package main

type Server struct {
	// an interface that the server needs to have
	gRPC.UnimplementedTemplateServer

	// here you can impliment other fields that you want
}

func (s *Server) SendPost(ctx context.Context, Message *message) (*Message, error) {
    //some code here
    ...
    ack :=  // make an instance of your return type
    return (ack, nil)
}