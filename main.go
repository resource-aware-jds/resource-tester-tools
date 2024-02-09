package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

const (
	MBToReserve = 1024
)

func main() {
	a := make([]bool, MBToReserve*52377650)
	fmt.Println("Reserved: ", " : ", MBToReserve, "mb")

	for _, i := range a {
		fmt.Println(i)
		break
	}

	// Gracefully Shutdown
	// Make channel listen for signals from OS
	gracefulStop := make(chan os.Signal, 1)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)

	<-gracefulStop

}
