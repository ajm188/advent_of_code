package alu

type ALU struct {
	W, X, Y, Z int
	Input      []int
}

func (alu *ALU) Execute(instructions []Instruction) {
	for _, instr := range instructions {
		instr.Execute(alu)
	}
}
