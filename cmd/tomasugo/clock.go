package main

// provides the interface for ticking the clock

type Clock struct {
	Cycle int
}

func NewClock() *Clock {
	return &Clock{
		Cycle: 0,
	}
}

func (cl *Clock) Tick() {
	cl.Cycle = cl.Cycle + 1
}
