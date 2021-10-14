package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/dierbei/blind-box/initialize"
)

func main() {
	initialize.HttpServerRun()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	initialize.HttpServerStop()
}
