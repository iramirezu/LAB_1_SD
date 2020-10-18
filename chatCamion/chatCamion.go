package chatCamion

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) FuncHolaMUndo(ctx context.Context, mensaje *MensajeRequest) (*MensajeReply, error) {
	log.Printf("Mensaje desde Camion: %s", MensajeRequest.GetMensaje1())
	return &MensajeReply{Respuesta1: "Hola Camion qlo"}, nil
}