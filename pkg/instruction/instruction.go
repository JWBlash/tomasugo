package instruction

// instructions that flow through the algo

var InstructionQueue = make(chan Instruction)

type Instruction struct {
	Addr string
	Op   string
}
