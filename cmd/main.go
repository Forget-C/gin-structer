package main

import (
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/Forget-C/http-structer/internal/server"
	flag "github.com/spf13/pflag"
)

func main() {
	flag.Parse()
	defer func() {
		if r := recover(); r != nil {
			const size = 64 << 10
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
			log.Println("[main] recover painc:%s. stack:%s", r, string(buf))
			exit(-1)
		}
	}()

	go func() {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals,
			syscall.SIGTERM,
			syscall.SIGINT,
			syscall.SIGHUP,
			syscall.SIGQUIT,
		)
		<-signals
		exit(0)
	}()
	config.Parse()

	server.Run()
}

func exit(status int) {
	server.Stop()
	log.Println("[main] exit:%d.", status)
	os.Exit(status)
}
