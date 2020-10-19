package chatCliente

import (
	"log"
	"golang.org/x/net/context"
)

func (s *Server) GenerarOrden(ctx context.Context, mensaje *OrdenGenerada) (*IdSeguimiento, error) {
	IteradorIdSeguimiento = IteradorIdSeguimiento +1
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