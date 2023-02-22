package services

import (
	"context"
	"fmt"
)

type CalculatorService interface {
	Hello(name string) error
}

type calculatorService struct {
	calculatorClient CalculatorClient
}

func NewCalculatorService(calculatorClient CalculatorClient) CalculatorService {
	return calculatorService{calculatorClient}
}

func (base calculatorService) Hello(name string) error {
	req := HelloRequest{
		Name: name,
	}
	res, err := base.calculatorClient.Hello(context.Background(), &req)
	if err != nil {
		return err
	}
	
	fmt.Printf("Service: Hello")
	fmt.Printf("Request: %v\n", req.Name)
	fmt.Printf("Response: %v\n", res.Result)
	return nil 
}

