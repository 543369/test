package config

// Config 应用配置
type Config struct {
	Server ServerConfig
}

// ServerConfig 服务器配置
type ServerConfig struct {
	Port            string
	AllowedOrigins  []string
	AllowedMethods  []string
	AllowedHeaders  []string
}

// NewDefaultConfig 创建默认配置
func NewDefaultConfig() *Config {
	return &Config{
		Server: ServerConfig{
			Port:           "8080",
			AllowedOrigins: []string{"*"},
			AllowedMethods: []string{"GET", "POST", "OPTIONS"},
			AllowedHeaders: []string{"Content-Type", "Connect-Protocol-Version"},
		},
	}
}