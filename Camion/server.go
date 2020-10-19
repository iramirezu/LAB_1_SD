package main;

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/PrestigioExpress/ServicioCamion/chatCamion"
)

type server struct{}

func main()  {

	fmt.Println("Comenzando ejecucion de sistema Logistica-Camion")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 4040))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := chatCamion.Server{}

	grpcServer := grpc.NewServer()

	chatCamion.RegisterServicioCamionServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

	

}

