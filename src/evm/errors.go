package evm

import "fmt"

type UnsupportedOpcodeError struct {
	Opcode byte
}

func (e *UnsupportedOpcodeError) Error() string {
	return fmt.Sprintf("unsupported opcode 0x%02X", e.Opcode)
}

func NewUnsupportedOpcodeError(opcode byte) *UnsupportedOpcodeError {
	return &UnsupportedOpcodeError{
		Opcode: opcode,
	}
}
