package server

import (
	"calculator/api/proto/gen/genconnect"
	"calculator/internal/calculator"
	"calculator/internal/config"
	"calculator/pkg/logger"
	"fmt"
	"net/http"

	"github.com/bufbuild/connect-go"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

// Server 应用服务器
type Server struct {
	config *config.Config
	mux    *http.ServeMux
}

// NewServer 创建新服务器
func NewServer(cfg *config.Config) *Server {
	return &Server{
		config: cfg,
		mux:    http.NewServeMux(),
	}
}

// RegisterServices 注册所有服务
func (s *Server) RegisterServices() {
	// 创建计算器服务
	calculatorServer := &calculator.CalculatorServer{}
	
	// 创建拦截器
	interceptors := connect.WithInterceptors(LoggingInterceptor())
	
	// 注册计算器服务
	calculatorPath, calculatorHandler := genconnect.NewCalculatorServiceHandler(
		calculatorServer,
		interceptors,
	)
	s.mux.Handle(calculatorPath, calculatorHandler)
	
	logger.Info.Printf("已注册计算器服务于路径: %s", calculatorPath)
}

// Start 启动服务器
func (s *Server) Start() error {
	addr := fmt.Sprintf(":%s", s.config.Server.Port)
	
	// 添加CORS中间件
	corsMiddleware := CorsMiddleware(&s.config.Server)
	
	logger.Info.Printf("启动服务器在 %s 端口...", s.config.Server.Port)
	return http.ListenAndServe(
		addr,
		corsMiddleware(h2c.NewHandler(s.mux, &http2.Server{})),
	)
}