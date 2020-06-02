package share

import (
	"database/sql"
	"os"
)

// 日志文件
var ShareLoggerFile *os.File

// 全局 mysql DB 连接据柄
var ShareDb *sql.DB