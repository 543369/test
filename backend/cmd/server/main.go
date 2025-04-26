package main

import (
	"context"
	"log"
	"net/http"

	"github.com/bufbuild/connect-go"
	"calculator/internal/calculator"
	calculatorv1 "calculator/proto"
	calculatorconnect "calculator/proto/calculatorconnect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	calculatorServer := &calculator.CalculatorServer{}
	path, handler := calculatorconnect.NewCalculatorServiceHandler(
		calculatorServer,
		connect.WithInterceptors(
			// 添加CORS支持
			func(next connect.UnaryFunc) connect.UnaryFunc {
				return func(
					ctx context.Context,
					req connect.AnyRequest,
				) (connect.AnyResponse, error) {
					return next(ctx, req)
				}
			},
		),
	)

	mux := http.NewServeMux()
	mux.Handle(path, handler)

	// 添加CORS中间件
	corsHandler := func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Connect-Protocol-Version")

			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}

			h.ServeHTTP(w, r)
		})
	}

	log.Println("启动服务器在 :8080 端口...")
	http.ListenAndServe(
		":8080",
		corsHandler(h2c.NewHandler(mux, &http2.Server{})),
	)
}
