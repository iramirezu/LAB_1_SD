package main;

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/PrestigioExpress/ServicioCliente/chatCliente"
)

type server struct{}

func main()  {

	fmt.Println("Comenzando ejecucion de sistema Logistica-Cliente")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 4050))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := chatCamion.Server{}

	grpcServer := grpc.NewServer()

	chatCamion.RegisterServicioClienteServer(grpcServer, &s) // actualizar

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

	

}

