package main

import (
	"os"
)

func main() {
	//TODO init config
	//TODO init program (di)
	go WaitForTermination(Shutdown)
	//TODO run program
	os.Exit(0)
}
