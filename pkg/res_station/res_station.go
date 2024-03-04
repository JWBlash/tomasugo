package res_station

// reservation stations

import "github.com/jwblash/tomasugo/pkg/instruction"

type ResStation struct {
	Name  string
	Slots chan instruction.Instruction
}
