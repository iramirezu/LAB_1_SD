package chatCliente

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) FuncHolaMUndo(ctx context.Context, mensaje *MensajeRequest) (*MensajeReply, error) {
	log.Printf("Mensaje desde Camion: %s", mensaje.Mensaje1)
	return &MensajeReply{Respuesta1: "Hola Cliente qlo"}, nil
}