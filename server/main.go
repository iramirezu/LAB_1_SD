package chatCamion;

import (
	"main"
	"fmt"
	"log"
	"net"
	"golang.org/x/net/context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/PrestigioExpress/ServicioCamion/chatCamion"
)

type server struct{}

func main()  {

	fmt.Println("Comenzando ejecucion de sistema Logistica-Camion")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := chatCamion.Server{}

	grpcServer := grpc.NewServer()

	chat.RegisterServicioCamionServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

	

}

