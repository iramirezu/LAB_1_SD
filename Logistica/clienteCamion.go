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

	"github.com/PrestigioExpress/ServicioDistribuido/chatCamion"
)

// Se informa a la central en cuanto se completa una entrega, haya sido recibida o no

// pedir 
//	Tiempo de espera de segundo paquete cuando se esta en bodega
//	Tiempo de espera hasta realizar otro intento
//	Tiempo que demora el camion para llegar a un domicilio

// Fecha de entrega: 
//		- Fecha cuando se entrega paquede
//		- Si es 0 el paquete no se entrego al cliente, ya no quedan intentos

func CrearCamion(numCamion int){
	// Conexion con servidor Logistica "dist125"
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("dist125:5151", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se pudo conectar a servidor Logistica: %s", err)
	}
	defer conn.Close()

	c := chatCamion.NewServicioCamionClient(conn)

	// Definicion tipo de camion y su respectivo nombre
	nombreCamion := ""
	tipoC := "" // "0" si es Retail, "1" si es Normal
	if numCamion == 0 {
		nombreCamion = "Camion Retail 1"
		fmt.Println("Camion Retail 1")
		tipoC = "0"
	} else if numCamion == 1 {
		nombreCamion = "Camion Retail 2"
		fmt.Println("Camion Retail 2")
		tipoC := "0"
	} else{ 
		nombreCamion = "Camion Normal 1"
		fmt.Println("Camion Normal 1")
		tipoC := "1"
	}

	// Comienza Funcionamiento de Camion
	

	response_p, err_p := c.PedirPaquete(context.Background(), &chatCamion.PeticionPaquete{TipoCamion: tipoC})
	if err_p != nil {
		log.Fatalf("Error al llamar funcion PedirPaquete: %s", err_p)
	}
	id_p := response.Id
	tipo_p := response.Tipo
	valor_p := response.Valor
	origen_p := response.Origen
	destino_p := response.Destino
	intentos_p := response.Intentos
	fechaEntrega_p := response.FechaEntrega
	exito_p := response.Exito

	fmt.Println("Id Recibida desde Logistica: " + id_p + ", " + tipo_p + ", " + valor_p + ", " + origen_p + ", " + destino_p + ", " + intentos_p + ", " + fechaEntrega_p + ", " + exito_p)


	id_c := "response.Id"
	tipo_c := "response.Tipo"
	valor_c := "response.Valor"
	origen_c := "response.Origen"
	destino_c := "response.Destino"
	intentos_c := "response.Intentos"
	fechaEntrega_c := "response.FechaEntrega"
	exito_c := "response.Exito"
	response_c, err_c := c.CompletarEntrega(context.Background(), &chatCamion.PaqueteCompletado{Id:id_c, Tipo:tipo_c, Valor:valor_c, Origen:origen_c, Destino:destino_c,Intentos:intentos_c, FechaEntrega:fechaEntrega_c, Exito:exito_c})
	if err_c != nil {
		log.Fatalf("Error al llamar funcion CompletarEntrega: %s", err_c)
	}
	fmt.Println("Orden 'paquete entregado' recibida por Logistica " + response_c)

	
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


