package main

import (
	"github.com/e421083458/golang_common/lib"
	"github.com/liuzw3018/gateway_admin/router"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// 如果configPath为空，则从命令行中 '--config=../etc/prod/'中读取
	lib.InitModule("../etc/dev/", []string{"base", "mysql", "redis"})
	defer lib.Destroy()
	router.HttpServerRun()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	router.HttpServerStop()
}
