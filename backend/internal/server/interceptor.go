package server

import (
	"calculator/pkg/logger"
	"context"
	"time"

	"github.com/bufbuild/connect-go"
)

// LoggingInterceptor 日志拦截器
func LoggingInterceptor() connect.UnaryInterceptorFunc {
	return connect.UnaryInterceptorFunc(func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			start := time.Now()
			logger.Info.Printf("请求开始: %s", req.Spec().Procedure)
			
			res, err := next(ctx, req)
			
			duration := time.Since(start)
			if err != nil {
				logger.Error.Printf("请求失败: %s, 错误: %v, 耗时: %v", 
					req.Spec().Procedure, err, duration)
			} else {
				logger.Info.Printf("请求成功: %s, 耗时: %v", 
					req.Spec().Procedure, duration)
			}
			
			return res, err
		}
	})
}