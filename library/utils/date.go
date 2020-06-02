package utils

import "time"

// 获取系统当前时间
func GetNowData() string {
	return time.Now().Format("2006/01/02 15:04:05")
}