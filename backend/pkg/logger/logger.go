package logger

import (
	"log"
	"os"
)

var (
	// Info 信息日志
	Info *log.Logger
	// Error 错误日志
	Error *log.Logger
)

// Init 初始化日志
func Init() {
	Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}