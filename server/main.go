package main

import (
	"fmt"
	"lebrancconvas/gogrpc/server/services"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()	
	
	listener, err := net.Listen("tcp", ":9002")
	if err != nil {
		log.Fatal(err)
	}

	services.RegisterCalculatorServer(s, services.NewCalculatorServer())

	fmt.Println("gRPC server is running on port 9002")
	err = s.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}