package logger

import (
	"log"
	"os"
	"fmt"
	"table/library/utils"
	"table/share"
)

var (
	infoPrefix    string = "[Info] "    // 重要的信息
	warningPrefix string = "[Warning] " // 需要注意的信息
	errorPrefix   string = "[Error] "   // 致命错误
	debugPrefix   string = "[Debug] "   // 调试信息
)

func Info(contents string)  {
	info := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	info.Println(infoPrefix + contents)
	utils.FileWrite(share.ShareLoggerFile, fmt.Sprintln(utils.GetNowData(), infoPrefix, contents))
}

func Error(contents string)  {
	//error := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Llongfile)
	error := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	error.Fatal(errorPrefix + contents)
	utils.FileWrite(share.ShareLoggerFile, fmt.Sprintln(utils.GetNowData(), errorPrefix, contents))
}

func Debug(contents string)  {
	debug := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Llongfile)
	debug.Println(debugPrefix + contents)
	utils.FileWrite(share.ShareLoggerFile, fmt.Sprintln(utils.GetNowData(), debugPrefix, contents))
}

func Warning(contents string)  {
	warning := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	warning.Println(warningPrefix + contents)
	utils.FileWrite(share.ShareLoggerFile, fmt.Sprintln(utils.GetNowData(), warningPrefix, contents))
}