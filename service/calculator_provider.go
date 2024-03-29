package service

import (
	"context"

	"github.com/halationxiv/calculator/request"
)

type CalculatorProvider interface {
	// Addition function to add number
	Addition(ctx context.Context, request request.InputCommand) error
	// Substraction function to substract number
	Substraction(ctx context.Context, request request.InputCommand) error
	// Multiplication function to multiply number
	Multiplication(ctx context.Context, request request.InputCommand) error
	// Division function to divide number
	Division(ctx context.Context, request request.InputCommand) error
	// Cancel function to reset number to 0
	Cancel(ctx context.Context, request request.InputCommand) error
	// Exit function to exit from program
	Exit(ctx context.Context, request request.InputCommand) error
	// Absolute function to turn number into positive number
	Absolute(ctx context.Context, request request.InputCommand) error
	// Negative function to turn number into negative number
	Negative(ctx context.Context, request request.InputCommand) error
	// Squareroot function to return square root of the number
	Squareroot(ctx context.Context, request request.InputCommand) error
	// Square function to raise number to the power of 2
	Square(ctx context.Context, request request.InputCommand) error
	// Cubert function to return the cube root of the number
	Cubert(ctx context.Context, request request.InputCommand) error
	// Cube function to raise the number to the power of 3
	Cube(ctx context.Context, request request.InputCommand) error
	// Repeat run several commands previously
	Repeat(ctx context.Context, request request.InputCommand) error
	// GetCurrentNumber get current number
	GetCurrentNumber(ctx context.Context) float64
}
