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

	"github.com/PrestigioExpress/ServicioCliente/chatCamion"
)

// Se informa a la central en cuanto se completa una entrega, haya sido recibida o no

// pedir 
//	Tiempo de espera de segundo paquete cuando se esta en bodega
//	Tiempo de espera hasta realizar otro intento
//	Tiempo que demora el camion para llegar a un domicilio

// Fecha de entrega: 
//		- Fecha cuando se entrega paquede
//		- Si es 0 el paquete no se entrego al cliente, ya no quedan intentos

func CrearCamion(tipoCamion int){
	// Inicio de Creacion Cliente Camion (GRPC)
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
	conn, err := grpc.Dial("dist125:5151", grpc.WithInsecure())
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
	fmt.Println("MensajeReply desde Logistica: " + response.Respuesta1)
	// Termino de Creacion Cliente Camion (GRPC)

	// Inicio Loop para funcionamiento de Camiones
	
}

func haySinEntregar() bool {
	rows := leerFilasRegistro();
	for i := 1; i < len(rows); i++{
		entrega := rows[i][6]
		if entrega == "0" {
			return true
		}
	}
	return false
}

func entregaLograda() bool {
	random := rand.Intn(100)
	if random < 19 { // < 20, pero funcion random empieza desde 0
		return false
	}else{
		return true
	}
}

func revisarIntentos(idPaquete string) int {
	rows := leerFilasRegistro();
	for i := 1; i < len(rows); i++{
		id := rows[i][0]
		if id == idPaquete {
			entero, err := strconv.Atoi(rows[i][5])
			if err != nil{
				return -1
			}
			return entero
		}
	}
	return -1
}

func agregarIntento(idPaquete string) int {
	rows := leerFilasRegistro();
	for i := 1; i < len(rows); i++{
		id := rows[i][0]
		if id == idPaquete {
			entero, err := strconv.Atoi(rows[i][5])
			if err != nil{
				return -1
			}
			entero = entero + 1
			s := strconv.Itoa(-42)
			rows[i][5] = s
			return entero
		}
	}
	return -1
}

func leerFilasRegistro() [][]string {
    f, err := os.Open("registroCamion.csv")
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


func writeChanges(rows [][]string) {
    f, err := os.Create("registroCamion2.csv")
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

	go CrearCamion(0)
	go CrearCamion(1)
	go CrearCamion(2)
	for {
		time.Sleep(5)
	}
}


