package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jwblash/tomasugo/pkg/clock"
)

func startClock() {
	cl := clock.CPUClock
	for {
		cl.Tick()
		fmt.Printf("+++ CYCLE: %v\n", cl.Cycle)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	terminate := make(chan os.Signal, 1)
	signal.Notify(terminate, syscall.SIGINT) // cleanup on interrupt
	go startClock()

	select {
	case t := <-terminate:
		fmt.Printf("+++ %v received. terminating...\n", t)
		// work done case later
	}

}
