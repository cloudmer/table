package main

import (
	"runtime"
	"table/service"
)

func main() {
	// 开启多核
	runtime.GOMAXPROCS(runtime.NumCPU())
	service.StartService()
}
