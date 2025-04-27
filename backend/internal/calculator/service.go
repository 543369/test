package calculator

import (
	"context"
	"errors"
	"fmt"

	gen "calculator/api/proto/gen"
	"github.com/bufbuild/connect-go"
)

// CalculatorServer 实现计算器服务
type CalculatorServer struct{}

// Calculate 执行计算操作
func (s *CalculatorServer) Calculate(
	ctx context.Context,
	req *connect.Request[gen.CalculateRequest],
) (*connect.Response[gen.CalculateResponse], error) {
	r := req.Msg
	var result float64
	
	// 验证操作符
	if err := validateOperator(r.Operator); err != nil {
		return nil, connect.NewError(
			connect.CodeInvalidArgument,
			err,
		)
	}
	
	// 执行计算
	result, err := calculate(r.FirstNumber, r.SecondNumber, r.Operator)
	if err != nil {
		return nil, connect.NewError(
			connect.CodeInvalidArgument,
			err,
		)
	}

	// 返回结果
	return connect.NewResponse(&gen.CalculateResponse{
		Result: result,
	}), nil
}

// validateOperator 验证操作符
func validateOperator(operator string) error {
	validOperators := map[string]bool{
		"+": true,
		"-": true,
		"*": true,
		"/": true,
	}
	
	if !validOperators[operator] {
		return fmt.Errorf("不支持的运算符: %s", operator)
	}
	
	return nil
}

// calculate 执行计算
func calculate(a, b float64, operator string) (float64, error) {
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("除数不能为零")
		}
		return a / b, nil
	default:
		// 这里不应该到达，因为已经验证过操作符
		return 0, fmt.Errorf("不支持的运算符: %s", operator)
	}
}