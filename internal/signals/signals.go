package signals

import (
	"fmt"
	"os"
	"syscall"
)

func CatchSig(ch chan os.Signal, done chan bool){
	sig := <- ch

	fmt.Println("sig recieved:", sig)

	switch sig{
	case syscall.SIGINT:
		fmt.Println("handling a SIGINT now!")
	case syscall.SIGTERM:
		fmt.Println("handling a SIGTERM in an entirely different way!")
	default:
		fmt.Println("unexpected signal recieved")
	}

	done <- true
}
