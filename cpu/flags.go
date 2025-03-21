package cpu

// TODO: create functions for work with flags	

type ConditionsCodes uint8

const (
	ZeroFlag ConditionsCodes = 1 << 0 // if the result of the operation is equal to zero
	ParityFlag ConditionsCodes = 1 << 2 // if the result of the operation is negative
	AuxCarryFlag ConditionsCodes = 1 << 4 // if the number of units as a result of the operation if even
	SignFlag ConditionsCodes = 1 << 7 // if during the execution of an arithmetic operation a carry to the next digit occurred
	ZeroFlag ConditionsCodes = 1 << 6 // Set if, when performing operations with BCD codes, a carry-over occurred from the 4th to the 5th digit
)