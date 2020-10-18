package main;

import (
	"log"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/PrestigioExpress/ServicioCamion/chatCamion"
)

func CrearCamion(tipoCamion int){
	nombreCamion := ""

	if tipoCamion == 0 {
		nombreCamion = "Camion Retail 1"
		fmt.Println("Camion Retail 1")
	} else if tipoCamion == 1 {
		nombreCamion = "Camion Retail 2"
		fmt.Println("Camion Retail 2")
	} else{ 
		nombreCamion = "Camion Normal 1"
		fmt.Println("Camion Normal 1")
	}
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":4040", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se pudo conectar a servidor Logistica: %s", err)
	}
	defer conn.Close()

	c := chatCamion.NewServicioCamionClient(conn)

	mensajeCamion := "Hola desde Camion: " + nombreCamion + ""
	response, err := c.FuncHolaMUndo(context.Background(), &chatCamion.MensajeRequest{Mensaje1: mensajeCamion})
	if err != nil {
		log.Fatalf("Error al llamar funcion FuncHolaMUndo: %s", err)
	}
	fmt.Println("MensajeReply desde Logistica: %s", response.Respuesta1)
}
func main() {

	go CrearCamion(0)
	go CrearCamion(1)
	go CrearCamion(2)

}