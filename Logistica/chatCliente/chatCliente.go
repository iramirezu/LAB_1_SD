package chatCliente

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

func (s *Server) GenerarOrden(ctx context.Context, mensaje *OrdenGenerada) (*IdSeguimiento, error) {
	id := mensaje.Id
	producto := mensaje.Producto
	valor := mensaje.Valor
	tienda := mensaje.Tienda
	destino := mensaje.Destino
	tipo := mensaje.Tipo

    //t := time.Now()
	//timestamp := t.Format("2006-01-02 15:04:05")
	timestamp := "2006-01-02 15:04:05"
	
	filasRegistro := leerFilasRegistro("registroLogistica")
	intSeguimiento := cantFilasRegistro("registroLogistica")
	strSeguimiento := strconv.Itoa(intSeguimiento)

	// timestamp,id-paquete,tipo,nombre,valor,origen,destino,seguimiento,intentos,fecha_entrega
	var nuevaFila[]string
	nuevaFila = append(nuevaFila, timestamp, id, tipo, producto, valor, tienda, destino, strSeguimiento, "0", "0")
	filasRegistro = append(filasRegistro, nuevaFila)
	escribirFilasRegistro("registroLogistica", filasRegistro)

	log.Printf("Nueva Orden Generada (id_seguimiento: %s) (tipo: %s)", strSeguimiento, tipo)
	return &IdSeguimiento{Id: strSeguimiento}, nil
}

func (s *Server) ConsultarOrden(ctx context.Context, mensaje *IdSeguimiento) (*MensajeReply, error) {
	log.Printf("Consulta desde Cliente a id de seguimiento: %s", mensaje.Id)
	return &MensajeReply{Respuesta1: "Hola Cliente qlo"}, nil
}


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