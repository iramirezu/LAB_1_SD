package main;

import (
	"encoding/csv"
	"os"

	"time"

	"math/rand"
	"strconv"

	"sync"

	"log"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/PrestigioExpress/ServicioDistribuido/chatCamion"
)

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

// Se informa a la central en cuanto se completa una entrega, haya sido recibida o no

// pedir 
//	Tiempo de espera de segundo paquete cuando se esta en bodega
//	Tiempo de espera hasta realizar otro intento
//	Tiempo que demora el camion para llegar a un domicilio

// Fecha de entrega: 
//		- Fecha cuando se entrega paquede
//		- Si es 0 el paquete no se entrego al cliente, ya no quedan intentos

func (scounter *SafeCounter) CrearCamion(numCamion int, key string){
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
	registroCamion := ""
	if numCamion == 0 {
		nombreCamion = "Camion Retail 1"
		fmt.Println("Camion Retail 1")
		tipoC = "0"
		registroCamion = "registroCamion1"
	} else if numCamion == 1 {
		nombreCamion = "Camion Retail 2"
		fmt.Println("Camion Retail 2")
		tipoC = "0"
		registroCamion = "registroCamion2"
	} else{ 
		nombreCamion = "Camion Normal 1"
		fmt.Println("Camion Normal 1")
		tipoC = "1"
		registroCamion = "registroCamion3"
	}

	// Comienza Funcionamiento de Camion
	
	// Pide Paquetes hasta entontrar una respuesta con id_p != "00"
	// Luego de encontrar el primer paquete espera TIEMPO y realiza la peticion una segunda vez
	// El camion puede partir sin un segundo paquete
	id_p := "0"
	tiempoEsperaSegundoPaquete := 10
	tiempoEsperaCamino := 10
	numPaquetes := 0

	// LOOP QUE ESPERA PAQUETES ANTES DE SALIR A ENTREGARLOS
	for id_p == "0" {
		scounter.mux.Lock()
		fmt.Println("Camion: "+nombreCamion+" Esperando Paquetes")
		time.Sleep(time.Second*2) // TIEMPO ESPERA SEGUDNO PAQUETE???? SI
		// FUNCION ENVIA PETICION DE PAQUETE Y RECIBE PAQUETE PARA GUARDARDO EN REGISTRO
		response_p, err_p := c.PedirPaquete(context.Background(), &chatCamion.PeticionPaquete{TipoCamion: tipoC})
		if err_p != nil {
			log.Fatalf("Error al llamar funcion PedirPaquete: %s", err_p)
		}
		id_p = response_p.Id
		tipo_p := response_p.Tipo
		valor_p := response_p.Valor
		origen_p := response_p.Origen
		destino_p := response_p.Destino
		intentos_p := response_p.Intentos
		fechaEntrega_p := response_p.FechaEntrega
		exito_p := response_p.Exito
		if id_p != "0"{
			// UN Paquete ha sido recibido Correctamente
			fmt.Println("Se encontro un PRIMER paquete")
			fmt.Println("Id Recibida por camion: "+nombreCamion+" desde Logistica: " + id_p + ", " + tipo_p + ", " + valor_p + ", " + origen_p + ", " + destino_p + ", " + intentos_p + ", " + fechaEntrega_p + ", " + exito_p)
			// ESCRIBIR PAQUETES A CSV DE CAMION
			filasRegistro := leerFilasRegistro(registroCamion)
			var nuevaFila[]string
			nuevaFila = append(nuevaFila, id_p, tipo_p, valor_p, origen_p, destino_p, intentos_p, fechaEntrega_p, exito_p)
			filasRegistro = append(filasRegistro, nuevaFila)
			escribirFilasRegistro(registroCamion, filasRegistro)

			fmt.Println("Esperando SEGUNDO Paquete")
			time.Sleep(time.Second*2) 

			response_p, err_p := c.PedirPaquete(context.Background(), &chatCamion.PeticionPaquete{TipoCamion: tipoC})
			if err_p != nil {
				log.Fatalf("Error al llamar funcion PedirPaquete: %s", err_p)
			}
			id_p = response_p.Id
			tipo_p := response_p.Tipo
			valor_p := response_p.Valor
			origen_p := response_p.Origen
			destino_p := response_p.Destino
			intentos_p := response_p.Intentos
			fechaEntrega_p := response_p.FechaEntrega
			exito_p := response_p.Exito
			if id_p == "0"{
				fmt.Println("No Se encontro un SEGUNDO paquete")
				numPaquetes = 1
			} else{
				// UN Paquete ha sido recibido Correctamente
				fmt.Println("Se encontro un SEGUNDO paquete")
				fmt.Println("Id Recibida por camion: "+nombreCamion+" desde Logistica: " + id_p + ", " + tipo_p + ", " + valor_p + ", " + origen_p + ", " + destino_p + ", " + intentos_p + ", " + fechaEntrega_p + ", " + exito_p)
				// ESCRIBIR PAQUETES A CSV DE CAMION
				filasRegistro := leerFilasRegistro(registroCamion)
				var nuevaFila[]string
				nuevaFila = append(nuevaFila, id_p, tipo_p, valor_p, origen_p, destino_p, intentos_p, fechaEntrega_p, exito_p)
				filasRegistro = append(filasRegistro, nuevaFila)
				escribirFilasRegistro(registroCamion, filasRegistro)

				fmt.Println("Camion Saliendo de Bodega...")
				numPaquetes = 2
			}
		}
		scounter.v[key]++
		scounter.mux.Unlock()
	}
	// COMIENZA LOOP DE ENTREGA
	for  numPaquetes > 0 {
		time.Sleep(time.Second*2) 
		numPaquetes = numPaquetes - 1
	}
	fmt.Println("Entregas COMPLETADAS volviendo a Bodega...")


	// Al partir el camion espera TIEMPO de entrega y luego ve si el paquete pudo ser ENTREGADO
	// 		Si el paquete no es entregado y existe un segundo paquete por entregar
	// 			Al partir el camion espera TIEMPO de entrega y luego ve si el paquete pudo ser ENTREGADO
	//		Si el paquete es entregado y existe un segundo paquete
	//			Al partir el camion espera TIEMPO de entrega y luego ve si el paquete pudo ser ENTREGADO
	//		Si el paquete es entregado y NO existe un segundo paquete
	//			Entregas Completadas Vuelve a BODEGA




	// FUNCION QUE ENVIA ORDEN COMPLETADA CON DATOS ACTUALIZADOS DEL PAQUETE
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
	fmt.Println("Paquete entregado, envio de datos a Logistica Completado" + response_c.Respuesta1)

	
}


// Actualiza fechaEntrega y Exito de paquete completado, retorna lista con nuevas filas
func completarEntrega(registro string, idPaquete string, exito string) [][]string {
	rows := leerFilasRegistro(registro);
	for i := 1; i < len(rows); i++{
		id := rows[i][0]
		if id == idPaquete {
			
			rows[i][7] = exito
			escribirFilasRegistro(registro, rows)
			return rows
		}
	}
	return nil
}

// Funcion con random que devuelve true un 80% de las veces (Usado para ver si un intento de entrega fue lograda o no)
func entregaLograda() bool {
	random := rand.Intn(100)
	if random < 19 { // < 20, pero funcion random empieza desde 0
		return false
		}else{
			return true
		}
	}
	
// Permite revisar si existen entregas
// devolver id paquete por entregar
func haySinEntregar(registro string) bool {
	rows := leerFilasRegistro(registro);
	for i := 1; i < len(rows); i++{
		entrega := rows[i][6] // revisa si fecha Entrega es igual a "0", si es asi quedan intentos de entrega
		if entrega == "0" {
			return true
		}
	}
	return false
}
// devuelve cantidad de intentos de un paquete
func revisarIntentos(registro string, idPaquete string) int {
	rows := leerFilasRegistro(registro);
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

// Agrega intento a Registro de Camion que tiene idPaquete
func agregarIntento(registro string, idPaquete string) int {
	rows := leerFilasRegistro(registro);
	for i := 1; i < len(rows); i++{
		id := rows[i][0]
		if id == idPaquete {
			entero, err := strconv.Atoi(rows[i][5])
			if err != nil{
				return -1
			}
			entero = entero + 1
			s := strconv.Itoa(entero)
			rows[i][5] = s
			escribirFilasRegistro(registro, rows)
			return entero
		}
	}
	return -1
}

func leerFilasRegistro(registro string) [][]string {
    f, err := os.Open("Registros/"+registro+".csv")
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

func eliminarFilaRegistro(nombreRegistro string, index int) {
	a := leerFilasRegistro(nombreRegistro)
	copy(a[index:], a[index+1:]) // Shift a[i+1:] left one index.
	a = a[:len(a)-1]     // Truncate slice.
	escribirFilasRegistro(nombreRegistro, a)
}

func main() {
	scounter := SafeCounter{v: make(map[string]int)}

	go scounter.CrearCamion(2,"2") // Camion Retail 3
	go scounter.CrearCamion(1,"1") // Camion Retail 2
	go scounter.CrearCamion(0,"0") // Camion Retail 1
	for {
		time.Sleep(time.Second*5)
	}
}


