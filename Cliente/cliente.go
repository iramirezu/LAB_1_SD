package main;

import (
	"encoding/csv"
	"os"

	"time"

	"math/rand"
	"strconv"

	"log"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/PrestigioExpress/ServicioCamion/chatCliente"
)


// se pide tiempo entre consultas de cliente

/* 
Cliente:
	- Crea cliente grpc conectado a servidor en sistema Logistica
	- Lee una orden de excel pymes o retail dependiendo del tipo de cliente
	- Borra la fila y luego la agrega a excel con consultas finalizadas (para no perder consultas y usarlas en siguientes simulaciones)
	- Realiza 3 consultas de seguimiendo a la orden antes de realizar otra orden

*/
func CrearCliente(tipoCliente string, tiempoOrden int){
	if tipoCliente == "retail" {
		tipoCliente = "retail"
	} else {
		tipoCliente = "pymes"
	}

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":4050", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se pudo conectar a servidor Logistica: %s", err)
	}
	defer conn.Close()

	c := chatCliente.NewServicioClienteClient(conn) // actualizar

	mensajeCliente := "Hola desde Cliente "
	response, err := c.FuncHolaMUndo(context.Background(), &chatCliente.MensajeRequest{Mensaje1: mensajeCliente}) // actualizar
	if err != nil {
		log.Fatalf("Error al llamar funcion FuncHolaMUndo: %s", err)
	}
	fmt.Println("MensajeReply desde Logistica: " + response.Respuesta1)

	// Inicio Loop para funcionamiento de Camiones
	for {
		time.Sleep(5)
	}
	rows := leerFilasRegistro(tipoCliente)
	for i := 1; i < len(rows); i++{
		fmt.Println(rows[1])
	}

}





func leerFilasRegistro(nombreRegistro string) [][]string {
    f, err := os.Open(""+nombreRegistro+".csv")
    if err != nil {
        log.Fatal(err)
    }
    rows, err := csv.NewReader(f).ReadAll()
    f.Close()
    if err != nil {
        log.Fatal(err)
    }
    return rows
}

func escribirFilasRegistro(nombreRegistro string, rows [][]string) {
    f, err := os.Create(""+nombreRegistro+".csv")
    if err != nil {
        log.Fatal(err)
    }
    err = csv.NewWriter(f).WriteAll(rows)
    f.Close()
    if err != nil {
        log.Fatal(err)
    }
}


func main() {

	tiempoOrden := 1
	go CrearCamion("pymes", tiempoOrden)
	go CrearCamion("retail", tiempoOrden)
	for {
		time.Sleep(5)
	}
}


