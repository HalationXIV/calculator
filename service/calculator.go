package service

import (
	"context"
	"math"
	"os"

	"github.com/halationxiv/calculator/constant"
	"github.com/halationxiv/calculator/errors"
	"github.com/halationxiv/calculator/request"
)

type calculator struct {
	number      float64
	commandList []request.InputCommand
}

// NewCalculator initialize calculator that implements CalculatorProvider
func NewCalculator() CalculatorProvider {
	return &calculator{commandList: []request.InputCommand{}}
}

func (c *calculator) Addition(ctx context.Context, request request.InputCommand) error {
	c.number += float64(request.Value)
	c.displayResult(ctx, request)
	return nil
}
func (c *calculator) Substraction(ctx context.Context, request request.InputCommand) error {
	c.number -= float64(request.Value)
	c.displayResult(ctx, request)
	return nil
}
func (c *calculator) Multiplication(ctx context.Context, request request.InputCommand) error {
	c.number *= float64(request.Value)
	c.displayResult(ctx, request)
	return nil
}
func (c *calculator) Division(ctx context.Context, request request.InputCommand) error {
	if request.Value == 0 {
		return errors.Error_invalid_division_by_zero
	}
	c.number /= float64(request.Value)
	c.displayResult(ctx, request)
	return nil
}
func (c *calculator) Cancel(ctx context.Context, request request.InputCommand) error {
	c.number = 0
	c.displayResult(ctx, request)
	return nil
}
func (c *calculator) Exit(ctx context.Context, request request.InputCommand) error {
	os.Exit(1)
	return nil
}
func (c *calculator) Absolute(ctx context.Context, request request.InputCommand) error {
	if c.number < 0 {
		c.number *= -1
	}
	c.displayResult(ctx, request)
	return nil
}
func (c *calculator) Negative(ctx context.Context, request request.InputCommand) error {
	if c.number > 0 {
		c.number *= -1
	}
	c.displayResult(ctx, request)
	return nil
}
func (c *calculator) Squareroot(ctx context.Context, request request.InputCommand) error {
	c.number = math.Sqrt(c.number)
	c.displayResult(ctx, request)
	return nil
}
func (c *calculator) Square(ctx context.Context, request request.InputCommand) error {
	c.number *= c.number
	c.displayResult(ctx, request)
	return nil
}
func (c *calculator) Cubert(ctx context.Context, request request.InputCommand) error {
	c.number = math.Cbrt(c.number)
	c.displayResult(ctx, request)
	return nil
}
func (c *calculator) Cube(ctx context.Context, request request.InputCommand) error {
	c.number = c.number * c.number * c.number
	c.displayResult(ctx, request)
	return nil
}
func (c *calculator) Repeat(ctx context.Context, req request.InputCommand) error {
	if req.Value > constant.MAX_COMMAND_NUMBER {
		return errors.Error_invalid_repeat_number_input_more_than_max
	}

	if req.Value > len(c.commandList) {
		return errors.Error_invalid_repeat_number_input_more_than_array_limit
	}

	for i := len(c.commandList) - req.Value; i < len(c.commandList); i++ {
		reqI := c.commandList[i]
		reqI.Source = constant.SOURCE_REPEAT_COMMAND

		if fun, ok := MAP_CMD[reqI.Cmd]; ok {
			fun.(func(CalculatorProvider, context.Context, request.InputCommand) error)(c, ctx, reqI)
		}
	}

	return nil
}

func (c *calculator) GetCurrentNumber(ctx context.Context) float64 {
	return c.number
}
