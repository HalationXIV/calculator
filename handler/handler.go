package handler

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/halationxiv/calculator/constant"
	"github.com/halationxiv/calculator/errors"
	"github.com/halationxiv/calculator/request"
	"github.com/halationxiv/calculator/service"
)

type calculatorHandler struct {
	calc service.CalculatorProvider
}

// NewCalculatorHandler initialize calculator handler
func NewCalculatorHandler(calc service.CalculatorProvider) *calculatorHandler {
	return &calculatorHandler{calc: calc}
}

// handleCmd run main service function for calculator
func (h *calculatorHandler) handleCmd(ctx context.Context, line string) error {
	fields := strings.Fields(line)
	in := h.parseInput(ctx, fields)
	if in.Err != nil {
		return in.Err
	}
	if fun, ok := service.MAP_CMD[fields[0]]; ok {
		fun.(func(service.CalculatorProvider, context.Context, request.InputCommand) error)(h.calc, ctx, in)
	}

	return nil
}

// parseInput parse the input from command line and check validation
func (h *calculatorHandler) parseInput(ctx context.Context, fields []string) request.InputCommand {
	var in request.InputCommand
	in.Source = constant.SOURCE_USER_INPUT

	if len(fields) == 0 || len(fields) > 2 {
		in.Err = errors.Error_invalid_command
		return in
	}

	if _, ok := service.MAP_CMD[fields[0]]; ok {
		cmd := fields[0]
		switch cmd {
		case constant.CMD_ADDITION, constant.CMD_SUBSTRACTION, constant.CMD_MULTIPLICATION, constant.CMD_DIVISION, constant.CMD_REPEAT:
			if len(fields) == 1 {
				in.Err = errors.Error_invalid_parameter_for_command
				return in
			}
			value := fields[1]
			valInt, err := strconv.Atoi(value)
			if err != nil {
				in.Err = err
				return in
			}
			in.Cmd = cmd
			in.Value = valInt
		default:
			in.Cmd = cmd
		}
	} else {
		in.Err = errors.Error_invalid_command
	}

	return in
}

// HandleScanning handle scanning from user command line
func (h *calculatorHandler) HandleScanning(scanner *bufio.Scanner, calc service.CalculatorProvider) error {
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		ctx := context.Background()
		err := h.handleCmd(ctx, line)
		if err != nil {
			fmt.Println(err)
		}
	}
	if err := scanner.Err(); err != nil {
		if err != io.EOF {
			return err
		}
	}

	return nil
}
