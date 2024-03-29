package service

import "github.com/halationxiv/calculator/constant"

var (
	// MAP_CMD map the command the respective function of provider
	MAP_CMD = map[string]interface{}{
		constant.CMD_ADDITION:       CalculatorProvider.Addition,
		constant.CMD_SUBSTRACTION:   CalculatorProvider.Substraction,
		constant.CMD_MULTIPLICATION: CalculatorProvider.Multiplication,
		constant.CMD_DIVISION:       CalculatorProvider.Division,
		constant.CMD_SQUARE_ROOT:    CalculatorProvider.Squareroot,
		constant.CMD_SQUARE:         CalculatorProvider.Square,
		constant.CMD_CUBERT:         CalculatorProvider.Cubert,
		constant.CMD_CUBE:           CalculatorProvider.Cube,
		constant.CMD_ABSOLUTE:       CalculatorProvider.Absolute,
		constant.CMD_NEGATIVE:       CalculatorProvider.Negative,
		constant.CMD_REPEAT:         CalculatorProvider.Repeat,
		constant.CMD_CANCEL:         CalculatorProvider.Cancel,
		constant.CMD_EXIT:           CalculatorProvider.Exit,
	}
)
