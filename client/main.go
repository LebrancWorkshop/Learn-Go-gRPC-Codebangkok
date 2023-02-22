package main

import (
	"lebrancconvas/gogrpc/client/services"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	creds := insecure.NewCredentials()
	cc, err := grpc.Dial("localhost:9002", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()
	
	calculatorClient := services.NewCalculatorClient(cc)
	calculatorService := services.NewCalculatorService(calculatorClient)

	err = calculatorService.Hello("Nagi")
	if err != nil {
		log.Fatal(err)
	}
}