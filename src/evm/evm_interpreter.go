package evm

import (
	"strconv"
)

type EvmInterpreter struct {
	Stack        *Stack
	bytecode     string
	currentIndex int
}

func NewEvmInterpreter(bytecode string) *EvmInterpreter {
	return &EvmInterpreter{
		Stack:        NewStack(0),
		bytecode:     bytecode,
		currentIndex: 0,
	}
}

const (
	OpAdd   byte = 0x01
	OpPush1 byte = 0x60
	OpSwap1 byte = 0x90
)

var Opcodes = map[byte]Opcode{
	OpAdd:   Add{},
	OpPush1: Push1{},
	OpSwap1: Swap1{},
}

func (ec *EvmInterpreter) execute() error {
	// Advance after 0x prefix
	if ec.bytecode[:2] == "0x" {
		ec.next()
	}

	for {
		if !ec.hasNext() {
			break
		}

		next := ec.next()
		opcode, err := lookupOpcode(next)
		if err != nil {
			return err
		}

		err = opcode.Execute(ec)
		if err != nil {
			return err
		}
	}
	return nil
}

func lookupOpcode(hex string) (Opcode, error) {
	parsed, err := strconv.ParseInt(hex, 16, 16)
	if err != nil {
		return nil, err
	}

	opcodeByte := byte(parsed)
	opcode, exists := Opcodes[opcodeByte]
	if !exists {
		return nil, NewUnsupportedOpcodeError(opcodeByte)
	}
	return opcode, nil
}

// Returns next hex value in the bytecode
func (ec *EvmInterpreter) next() string {
	next := ec.bytecode[ec.currentIndex : ec.currentIndex+2]
	ec.currentIndex = ec.currentIndex + 2
	return next
}

func (ec *EvmInterpreter) hasNext() bool {
	return ec.currentIndex < len(ec.bytecode)-1
}
