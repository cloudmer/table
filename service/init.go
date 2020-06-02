package service

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"table/library/logger"
	"table/library/utils"
	"table/share"
	"time"
)

var op Option

func init()  {
	// 验证命令行参数
	FlagChecking()
	// 获取配置项
	op.GetOption()
	// runtime 文件夹创建
	runtimeDir()
	// 创建 程序 logger 文件
	loggerFile()
	// 打开Mysql
	connectMysql()
}

// runtime 文件夹
func runtimeDir()  {
	// 如果没有 配置 runtime 文件夹 就在当前进程目录下创建 runtime 文件夹
	if op.RuntimeDir == "" {
		pwd, _ := os.Getwd()
		op.RuntimeDir = pwd + "/runtime"
	}
	_, err := utils.GenerateDir(op.RuntimeDir)
	if err != nil {
		fmt.Println("runtime 文件夹创建失败, 请手动创建并赋予可写可读权限", err.Error())
		os.Exit(0)
	}
}

// logger 文件
func loggerFile()  {
	// 如果有 logger 文件 则追究 如果没有 logger 则 创建
	_, err := utils.GenerateFile(op.RuntimeDir +"/"+ op.LoggerFileName)
	if err != nil {
		fmt.Println("logger 文件创建失败", err.Error())
		os.Exit(0)
	}
}

func connectMysql()  {
	// 参数坚持
	if op.Mysql.Host == "" {
		log.Println("请先配置 mysql host")
		os.Exit(0)
	}
	if op.Mysql.Port == "" {
		log.Println("请先配置 mysql port")
		os.Exit(0)
	}
	if op.Mysql.Databases == "" {
		log.Println("请先配置 mysql databases")
		os.Exit(0)
	}
	if op.Mysql.Username == "" {
		log.Println("请先配置 mysql username")
		os.Exit(0)
	}
	if op.Mysql.Password == "" {
		log.Println("请先配置 mysql Password")
		os.Exit(0)
	}
	if op.Mysql.Charset == "" {
		log.Println("请先配置 mysql charset")
		os.Exit(0)
	}

	//dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", op.Mysql.Username, op.Mysql.Password, op.Mysql.Host, op.Mysql.Port, op.Mysql.Databases, op.Mysql.Charset)
	// 可以执行多条sql
	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=true", op.Mysql.Username, op.Mysql.Password, op.Mysql.Host, op.Mysql.Port, op.Mysql.Databases)
	fmt.Println(dbDSN)
	db, err := sql.Open("mysql", dbDSN)
	if err != nil {
		logger.Warning(err.Error())
		os.Exit(0)
	}

	// 测试连接
	err = db.Ping()
	if err != nil {
		logger.Warning(err.Error())
		os.Exit(0)
	}

	if op.Mysql.MaxOpen != 0 {
		// 最大连接数
		db.SetMaxOpenConns(op.Mysql.MaxOpen)
	}

	if op.Mysql.MaxIdle != 0 {
		// 闲置连接数
		db.SetMaxIdleConns(op.Mysql.MaxIdle)
	}

	if op.Mysql.MaxLifetime != 0 {
		// 最大连接周期
		db.SetConnMaxLifetime(time.Duration(op.Mysql.MaxLifetime) * time.Second)
	}

	logger.Info("mysql 连接成功")
	share.ShareDb = db
}
