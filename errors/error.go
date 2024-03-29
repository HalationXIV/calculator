package errors

import "errors"

var (
	Error_invalid_command                                   = errors.New("invalid command")
	Error_invalid_division_by_zero                          = errors.New("invalid division by zero")
	Error_invalid_repeat_number_input_more_than_array_limit = errors.New("invalid repeat command number, maximum %d command")
	Error_invalid_repeat_number_input_more_than_max         = errors.New("invalid repeat command number, maximum 10 command")
	Error_invalid_parameter_for_command                     = errors.New("invalid parameter for %s command")
)
