package evm

import (
	"errors"
	"testing"
)

func TestBytecodeInterpreter(t *testing.T) {
	bytecode := "0x6001600101"

	interpreter := EvmInterpreter{
		Stack:        NewStack(0),
		bytecode:     bytecode,
		currentIndex: 0,
	}

	err := interpreter.execute()
	if err != nil {
		t.Error(err)
	}

	if interpreter.Stack.Size() != 1 {
		t.Error("expected 1 element after execution")
	}
}

func TestReturnsErrorWhenOpcodeIsNotSupported(t *testing.T) {
	bytecode := "0x6101"

	interpreter := EvmInterpreter{
		Stack:        NewStack(0),
		bytecode:     bytecode,
		currentIndex: 0,
	}

	err := interpreter.execute()

	if err == nil {
		t.Error("expected UnsupportedOpcodeError")
	} else {
		var unsupportedErr *UnsupportedOpcodeError
		if !errors.As(err, &unsupportedErr) {
			t.Error("expected UnsupportedOpcodeError")
		}
	}
}
