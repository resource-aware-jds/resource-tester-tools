package main

import (
	"crypto/sha256"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	s := "sha256 this string"
	fmt.Println("Producting constant work load")
	for j := 0; j < 7; j++ {
		go func() {
			for {
				for i := 0; i < 1000000; i++ {
					h := sha256.New()
					h.Write([]byte(s))
					h.Sum(nil)
				}
				time.Sleep(1 * time.Millisecond)
			}
		}()
	}

	// Gracefully Shutdown
	// Make channel listen for signals from OS
	gracefulStop := make(chan os.Signal, 1)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)

	<-gracefulStop

}
