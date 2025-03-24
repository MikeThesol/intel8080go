package cpu

import "fmt"

type CPU struct {
	A      uint8
	B      uint8
	C      uint8
	D      uint8
	E      uint8
	H      uint8
	L      uint8
	PC     uint16 // program counter
	SP     uint16 // stack pointer
	Flags  ConditionsCodes
	Memory []uint8
}

func (cpu *CPU) executeInstraction() {
	opcode := cpu.Memory[cpu.PC]
	cpu.PC++

	switch opcode {
	case 0x00:
		//NOP
	case 0x01:
		//LXI B, #

	}
	// TODO: finish comands
}

func (cpu *CPU) Run() {
	fmt.Println("intel 8080 start")
	for {
		cpu.executeInstraction()
	}
}
