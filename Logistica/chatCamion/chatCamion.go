package chatCamion

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct {
}

import (
	"log"
	"encoding/csv"
	"strconv"
	"os"
	"golang.org/x/net/context"

	//"time"
)

type Server struct {
}


// =========================================== FUNCIONES DE GRPC =======================================================================================

func (s *Server) PedirPaquete(ctx context.Context, mensaje *PeticionPaquete) (*PaqueteRecibido, error) {
	tipoCamion := mensaje.TipoCamion

	// Falta leer datos de las Colas 
	// Leer infor de registro Logistica
	id_r := "id";
    tipo_r := "tipo";
    valor_r := "valor";
    origen_r := "origen";
    destino_r := "destino";
    intentos_r := "intentos";
    fechaEntrega_r := "fechaEntrega"; 
    exito_r := "exito";
    fmt.Println("Peticion de paquete desde camion: " + tipoCamion)
	
	return &PaqueteRecibido{Id: id_r, Tipo:tipo_r, Valor:valor_r, Origen:origen_r, Destino:destino_r, Intentos:intentos_r, FechaEntrega:fechaEntrega_r, Exito:exito_r}, nil
}

func (s *Server) CompletarEntrega(ctx context.Context, mensaje *PaqueteCompletado) (*MensajeReply, error) {
	id_c := mensaje.Id
	tipo_c := mensaje.Tipo
	valor_c := mensaje.Valor
	origen_c := mensaje.Origen
	destino_c := mensaje.Destino
	intentos_c := mensaje.Intentos
	fechaEntrega_c := mensaje.FechaEntrega
	exito_c := mensaje.Exito

	fmt.Println("Entrega Completada recibida desde camion: " + id_c + ", " + tipo_c + ", " + valor_c + ", " + origen_c + ", " + destino_c + ", " + intentos_c + ", " + fechaEntrega_c + ", " + exito_c)
	// escribir linea a Registro Logistica
	return &MensajeReply{respuesta1: "ta bien"}, nil
}





// =========================================== FUNCIONES DE AYUDA =======================================================================================

func cantFilasRegistro(nombreRegistro string) int { // numero de seguimiento sera un autogenerado que se relaciona directamente con la cantidad de registros
    rows := leerFilasRegistro(nombreRegistro)
	cant := (len(rows))	
    return cant
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

func eliminarFilaRegistro(nombreRegistro string, index int) {
	a := leerFilasRegistro(nombreRegistro)
	copy(a[index:], a[index+1:]) // Shift a[i+1:] left one index.
     // Erase last element (write zero value).
	a = a[:len(a)-1]     // Truncate slice.
	escribirFilasRegistro(nombreRegistro, a)
}