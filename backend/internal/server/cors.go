package server

import (
	"calculator/internal/config"
	"net/http"
	"strings"
)

// CorsMiddleware 创建CORS中间件
func CorsMiddleware(cfg *config.ServerConfig) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// 设置CORS头
			for _, origin := range cfg.AllowedOrigins {
				w.Header().Set("Access-Control-Allow-Origin", origin)
			}
			
			w.Header().Set("Access-Control-Allow-Methods", 
				strings.Join(cfg.AllowedMethods, ", "))
			
			w.Header().Set("Access-Control-Allow-Headers", 
				strings.Join(cfg.AllowedHeaders, ", "))

			// 处理预检请求
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}