package main

func main() {
	cpu := &CPU{
		Memory: make([]uint8, 65536),
	}
	cpu.Run()
}