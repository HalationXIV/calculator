package service

import (
	"context"
	"testing"

	"github.com/halationxiv/calculator/constant"
	"github.com/halationxiv/calculator/request"
	"github.com/stretchr/testify/assert"
)

func Test_calculator_Calculate_Sequence(t *testing.T) {
	ctx := context.Background()
	cTest := NewCalculator()

	// Test case sequence
	// add 2, neg, sqr, multiply 2, repeat 2, cancel

	cTest.Addition(ctx, request.InputCommand{
		Cmd:   constant.CMD_ADDITION,
		Value: 2,
	})
	cTest.Negative(ctx, request.InputCommand{
		Cmd:   constant.CMD_NEGATIVE,
		Value: 0,
	})
	cTest.Square(ctx, request.InputCommand{
		Cmd:   constant.CMD_SQUARE,
		Value: 0,
	})
	cTest.Multiplication(ctx, request.InputCommand{
		Cmd:   constant.CMD_MULTIPLICATION,
		Value: 2,
	})
	cTest.Repeat(ctx, request.InputCommand{
		Cmd:   constant.CMD_REPEAT,
		Value: 2,
	})
	assert.Equal(t, float64(128), cTest.GetCurrentNumber(ctx))

	// cancel, add 128, divide 2, cube, cubert, abs
	cTest.Cancel(ctx, request.InputCommand{
		Cmd:   constant.CMD_CANCEL,
		Value: 0,
	})
	cTest.Addition(ctx, request.InputCommand{
		Cmd:   constant.CMD_ADDITION,
		Value: 35,
	})
	cTest.Division(ctx, request.InputCommand{
		Cmd:   constant.CMD_DIVISION,
		Value: 2,
	})
	cTest.Cube(ctx, request.InputCommand{
		Cmd:   constant.CMD_CUBE,
		Value: 0,
	})
	cTest.Cubert(ctx, request.InputCommand{
		Cmd:   constant.CMD_CUBERT,
		Value: 0,
	})
	assert.Equal(t, float64(17.5), cTest.GetCurrentNumber(ctx))

	// cancel, addition 99, sqrt, substract 10, cubert, absolute, division 0
	cTest.Cancel(ctx, request.InputCommand{
		Cmd:   constant.CMD_CANCEL,
		Value: 0,
	})
	cTest.Addition(ctx, request.InputCommand{
		Cmd:   constant.CMD_ADDITION,
		Value: 99,
	})
	cTest.Squareroot(ctx, request.InputCommand{
		Cmd:   constant.CMD_SQUARE_ROOT,
		Value: 0,
	})
	cTest.Substraction(ctx, request.InputCommand{
		Cmd:   constant.CMD_SUBSTRACTION,
		Value: 10,
	})
	cTest.Cubert(ctx, request.InputCommand{
		Cmd:   constant.CMD_CUBERT,
		Value: 0,
	})
	cTest.Absolute(ctx, request.InputCommand{
		Cmd:   constant.CMD_ABSOLUTE,
		Value: 0,
	})
	cTest.Division(ctx, request.InputCommand{
		Cmd:   constant.CMD_DIVISION,
		Value: 0,
	})

	assert.Equal(t, float64(0.368711439107206), cTest.GetCurrentNumber(ctx))

	// cancel, substraction 99, addition 10, multiplication 2
	cTest.Cancel(ctx, request.InputCommand{
		Cmd:   constant.CMD_CANCEL,
		Value: 0,
	})
	cTest.Substraction(ctx, request.InputCommand{
		Cmd:   constant.CMD_SUBSTRACTION,
		Value: 99,
	})
	cTest.Addition(ctx, request.InputCommand{
		Cmd:   constant.CMD_ADDITION,
		Value: 10,
	})
	cTest.Multiplication(ctx, request.InputCommand{
		Cmd:   constant.CMD_MULTIPLICATION,
		Value: 2,
	})

	assert.Equal(t, float64(-178), cTest.GetCurrentNumber(ctx))

}
