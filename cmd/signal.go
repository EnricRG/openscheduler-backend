package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func WaitForTermination(onShutdown func(os.Signal)) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	signal := <-quit
	fmt.Printf("Received signal '%v', performing application shutdown", signal)
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Panic during application shutdown: %v", r)
		}
	}()
	onShutdown(signal)
}
