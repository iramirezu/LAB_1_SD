package chatCliente

import (
	"log"
	"strconv"
	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) GenerarOrden(ctx context.Context, mensaje *OrdenGenerada) (*IdSeguimiento, error) {
	IteradorIdSeguimiento := 1
	id := mensaje.id
	producto := mensaje.producto
	valor := mensaje.valor
	tienda := mensaje.tienda
	destino := mensaje.destino
	prioritario := mensaje.prioritario

	log.Printf("Nueva Orden Generada con id de producto: %s", id)
	
	filasRegistro := leerFilasRegistro("registroLogistica")
	var nuevaFila[]string
	nuevaFila = append(nuevaFila, id, producto, valor, tienda, destino, prioritario, strSeguimiento)
	filasRegistro = append(filasRegistro, nuevaFila)
	escribirFilasRegistro("registroLogistica", filasRegistro)
	strSeguimiento := strconv.Itoa(IteradorIdSeguimiento)
	return &IdSeguimiento{id: strSeguimiento}, nil
}

func (s *Server) ConsultarOrden(ctx context.Context, mensaje *IdSeguimiento) (*MensajeReply, error) {
	log.Printf("Consulta desde Cliente a id de seguimiento: %s", mensaje.id)
	return &MensajeReply{Respuesta1: "Hola Cliente qlo"}, nil
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