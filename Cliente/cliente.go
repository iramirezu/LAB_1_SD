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
	conn, err := grpc.Dial(":4072", grpc.WithInsecure())
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
		if rows != nil {
			// saca datos de primera linea en el csv que contiene intrucciones del cliente
			// elimina la linea y la agrega a "(pymes/retail)_usados" para que se puedan volver a usar las intrucciones
			id_r := rows[1][0] 
			producto_r := rows[1][1]
			valor_r := rows[1][2]
			tienda_r := rows[1][3]
			destino_r := rows[1][4]
	
			eliminarFilaRegistro(tipoCliente, 1)
	
			if tipoCliente != "retail" {
				prioritario := rows[1][5]
				if prioritario == "1" {
					tipoCliente = "prioritario"
				}else{
					tipoCliente = "normal"
				}
			}

			response, err := c.GenerarOrden(context.Background(), &chatCliente.OrdenGenerada{Id: id_r, Producto: producto_r, Valor:valor_r, Tienda:tienda_r, Destino:destino_r,Tipo:tipo_r}) // actualizar
			if err != nil {
				log.Fatalf("Error al llamar funcion GenerarOrden: %s", err)
			}
			fmt.Println("Id Seguimiento Generado: " + response.Id)
		}
	}

}





func leerFilasRegistro(nombreRegistro string) [][]string {
    f, err := os.Open("Registros/"+nombreRegistro+".csv")
    if err != nil {
		fmt.Println("No quedan mas registros 1")
        log.Fatal(err)
    }
    rows, err := csv.NewReader(f).ReadAll()
    f.Close()
    if err != nil {
		fmt.Println("No quedan mas registros 2")
		return nil
		//log.Fatal(err)
	}
	
	if len(rows) == 1 {
		return nil
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

func eliminarFilaRegistro(nombreRegistro string, index int) {
	a := leerFilasRegistro(nombreRegistro)
	copy(a[index:], a[index+1:]) // Shift a[i+1:] left one index.
	a = a[:len(a)-1]     // Truncate slice.
	escribirFilasRegistro(nombreRegistro, a)
}


func main() {

	tiempoOrden := 1
	go CrearCliente("retail", tiempoOrden)
	go CrearCliente("pymes", tiempoOrden)
	go CrearCliente("pymes", tiempoOrden)
	go CrearCliente("pymes", tiempoOrden)
	go CrearCliente("pymes", tiempoOrden)
	go CrearCliente("pymes", tiempoOrden)
	for {
		time.Sleep(5)
	}
}
