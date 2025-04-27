package main

import (
	"log"
	"net/http"

	"calculator/api/proto/gen/genconnect" // 你的 connect-go 生成的包
	"calculator/internal/calculator"
)

func main() {
	// 实例化你的服务实现
	calculatorService := &calculator.CalculatorServer{}

	// 创建 connect handler
	mux := http.NewServeMux()
	path, handler := genconnect.NewCalculatorServiceHandler(calculatorService)
	mux.Handle(path, handler)

	// 启动 HTTP 服务器，监听 8080 端口
	addr := ":8080"
	log.Printf("Calculator service listening on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
}
