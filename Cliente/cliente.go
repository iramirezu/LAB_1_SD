package main;

import (
	"encoding/csv"
	"os"

	"time"
	//"strconv"

	//"math/rand"

	"log"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/PrestigioExpress/ServicioCliente/chatCliente"
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
	conn, err := grpc.Dial(":4060", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se pudo conectar a servidor Logistica: %s", err)
	}
	defer conn.Close()

	c := chatCliente.NewServicioClienteClient(conn) // actualizar

	
	// Inicio Loop para funcionamiento de Clientes
	tipo_r := tipoCliente
	for {
		time.Sleep(5)
		rows := leerFilasRegistro(tipoCliente)
		//for i := 1; i < len(rows); i++{
		//	rows[i][2] = "0"
		//}
		// saca datos de primera linea en el csv que contiene intrucciones del cliente
		// elimina la linea y la agrega a "(pymes/retail)_usados" para que se puedan volver a usar las intrucciones
		id_r := rows[1][0] 
		producto_r := rows[1][1]
		valor_r := rows[1][2]
		tienda_r := rows[1][3]
		destino_r := rows[1][4]

		response, err := c.GenerarOrden(context.Background(), &chatCliente.OrdenGenerada{id_r,producto_r,valor_r,tienda_r,destino_r,tipo_r}) // actualizar
		if err != nil {
			log.Fatalf("Error al llamar funcion GenerarOrden: %s", err)
		}
		fmt.Println("Id Seguimiento Generado: " + response.Id)
	}

}





func leerFilasRegistro(nombreRegistro string) [][]string {
    f, err := os.Open("Registros/"+nombreRegistro+".csv")
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
    f, err := os.Create("Registros/"+nombreRegistro+".csv")
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
	go CrearCliente("pymes", tiempoOrden)
	go CrearCliente("retail", tiempoOrden)
	for {
		time.Sleep(5)
	}
}


