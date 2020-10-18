package proto;

import (
	"LAB_1_SD/proto"
	
	"log"
	"net"
	"golang.org/x/net/context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func main()  {
	listener, err := net.Listen("tcp" , ":4040")
	if err != nil {
		panic(err)
	}	

	srv := grpc.NewServer()
	proto.RegisterAddServiceServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}


}

