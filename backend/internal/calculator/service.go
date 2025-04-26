package calculator

import (
	"context"
	"errors"
	"fmt"

	"github.com/bufbuild/connect-go"
	calculatorv1 "calculator/proto"
)

// CalculatorServer 实现计算器服务
type CalculatorServer struct{}

// Calculate 执行计算操作
func (s *CalculatorServer) Calculate(
	ctx context.Context,
	req *connect.Request[calculatorv1.CalculateRequest],
) (*connect.Response[calculatorv1.CalculateResponse], error) {
	r := req.Msg
	var result float64
	var err error

	switch r.Operator {
	case "+":
		result = r.FirstNumber + r.SecondNumber
	case "-":
		result = r.FirstNumber - r.SecondNumber
	case "*":
		result = r.FirstNumber * r.SecondNumber
	case "/":
		if r.SecondNumber == 0 {
			return nil, connect.NewError(
				connect.CodeInvalidArgument,
				errors.New("除数不能为零"),
			)
		}
		result = r.FirstNumber / r.SecondNumber
	default:
		return nil, connect.NewError(
			connect.CodeInvalidArgument,
			fmt.Errorf("不支持的运算符: %s", r.Operator),
		)
	}

	return connect.NewResponse(&calculatorv1.CalculateResponse{
		Result: result,
	}), nil
}