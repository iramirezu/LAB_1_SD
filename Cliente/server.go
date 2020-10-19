package main;

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"golang.org/x/net/context"
	"github.com/PrestigioExpress/ServicioCliente/chatCliente"
)

type Server struct {
}

func (s *Server) FuncHolaMUndo(ctx context.Context, mensaje *MensajeRequest) (*MensajeReply, error) {
	log.Printf("Mensaje desde Camion: %s", mensaje.Mensaje1)
	return &MensajeReply{Respuesta1: "Hola Cliente qlo"}, nil
}

func main()  {

	fmt.Println("Comenzando ejecucion de sistema Logistica-Cliente")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 4050))
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

