package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/halationxiv/calculator/handler"
	"github.com/halationxiv/calculator/service"
)

func main() {
	// TODO: write unit test, test the code

	// Initialize Service and handler
	scanner := bufio.NewScanner(os.Stdin)
	c := service.NewCalculator()
	h := handler.NewCalculatorHandler(c)

	fmt.Println("Calculator service starting... ")

	err := h.HandleScanning(scanner, c)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error handleScanning %v", err))
	}

	fmt.Println("Calculator service Shutdown... ")

	os.Exit(1)
}
