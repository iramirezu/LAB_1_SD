package main;

import (
	"log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/PrestigioExpress/ServicioCamion/chatCamion"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":4040", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se pudo conectar a servidor Logistica: %s", err)
	}
	defer conn.Close()

	c := chat.NewServicioCamionClient(conn)

	response, err := c.FuncHolaMUndo(context.Background(), &chat.MensajeRequest{Mensaje1: "Hola desde Camion"})
	if err != nil {
		log.Fatalf("Error al llamar funcion FuncHolaMUndo: %s", err)
	}
	log.Printf("MensajeReply desde Logistica: %s", response.Respuesta1)

}