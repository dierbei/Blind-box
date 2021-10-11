package main

import (
	"github.com/dierbei/blind-box/initialize"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	initialize.HttpServerRun()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	initialize.HttpServerStop()
}
