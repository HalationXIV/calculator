package service

import (
	"context"
	"fmt"

	"github.com/halationxiv/calculator/constant"
	"github.com/halationxiv/calculator/request"
)

func (c *calculator) displayResult(ctx context.Context, request request.InputCommand) {
	if request.Source != constant.SOURCE_REPEAT_COMMAND {
		c.commandList = append(c.commandList, request)
	}
	fmt.Printf("%.1f\n", c.number)
}
