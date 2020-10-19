package main;

import (
	"fmt"
	"log"
	"net"

	//"encoding/csv"
	//"os"
	//"strconv"

	"google.golang.org/grpc"
	"github.com/PrestigioExpress/ServicioCliente/chatCliente"
)





func main() {
	fmt.Println("Comenzando ejecucion de sistema Logistica-Cliente")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 4072))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := chatCliente.Server{}

	grpcServer := grpc.NewServer()

	chatCliente.RegisterServicioClienteServer(grpcServer, &s) // actualizar

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

	

}



