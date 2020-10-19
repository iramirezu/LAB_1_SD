package main;

import (
	"fmt"
	"log"
	"net"

	"encoding/csv"
	"os"
	"strconv"

	"google.golang.org/grpc"
	"github.com/PrestigioExpress/ServicioCliente/chatCliente"
)

type Server struct {
}



func main() {
	IteradorIdSeguimiento := 0
	fmt.Println("Comenzando ejecucion de sistema Logistica-Cliente")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 4050))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := chatCliente.Server{}

	grpcServer := grpc.NewServer()

	chatCliente.RegisterServicioClienteServer(grpcServer, &s) // actualizar

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

	

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
