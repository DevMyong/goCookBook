package main

import (
	"fmt"
	"goServer/internal/signals"
	"os"
	"os/signal"
	"syscall"
)

func main(){
	signaltwo := make(chan os.Signal)
	done := make(chan bool)

	signal.Notify(signaltwo, syscall.SIGINT, syscall.SIGTERM)

	go signals.CatchSig(signaltwo, done)

	fmt.Println("Press ctrl-c to terminate...")
	<-done
	fmt.Println("Done!")
}