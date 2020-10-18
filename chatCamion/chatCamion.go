package main

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct {
}

func (s *server) FuncHolaMUndo(ctx context.Context, mensaje *MensajeRequest) (*MensajeReply, error) {
	log.Printtf("Mensaje desde Camion: %s", MensajeRequest.Mensaje1)
	return &MensajeReply{Respuesta1: "Hola Camion qlo"}, nil
}