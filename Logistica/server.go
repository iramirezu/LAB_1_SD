package main;

import (
	"fmt"
	"log"
	"net"

	//"encoding/csv"
	//"os"
	//"strconv"

	"time"

	"google.golang.org/grpc"
	"github.com/PrestigioExpress/ServicioCliente/chatCliente"
	"github.com/PrestigioExpress/ServicioCliente/chatCamion"
)



func serverChatCliente() {
	fmt.Println("Comenzando ejecucion de sistema Logistica-Cliente")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 5252))
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

func serverChatCamion() {
	fmt.Println("Comenzando ejecucion de sistema Logistica-Camion")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 5151))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := chatCamion.Server{}

	grpcServer := grpc.NewServer()

	chatCamion.RegisterServicioCamionServer(grpcServer, &s) // actualizar

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}


func main() {
	go serverChatCamion()
	go serverChatCliente()	
	for {
		time.Sleep(1)
	}
	

}



