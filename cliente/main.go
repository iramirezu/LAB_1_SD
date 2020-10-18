package main

import (
	"google.golang.org/grpc"
)

func main()  {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure()) // insecure porq no se usa https
	if err != nil {
		panic(err)
	}

	cliente := proto.NewAddServiceCliente(conn)

	g := gin.Default()


}