package main

import (
	"fmt"

	"github.com/jwblash/tomasugo/pkg/clock"
)

func main() {
	cl := clock.CPUClock
	fmt.Printf("%v\n", cl.Cycle)
	cl.Tick()
	fmt.Printf("%v\n", cl.Cycle)
}
