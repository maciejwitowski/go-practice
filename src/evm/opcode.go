package evm

import (
	"errors"
	"strconv"
)

type Opcode interface {
	Execute(ec *EvmInterpreter) error
}

type Add struct{}

func (add Add) Execute(ec *EvmInterpreter) error {
	if ec.Stack.Size() < 2 {
		return errors.New("not enough elements on the stack")
	}

	a, err := ec.Stack.pop()
	if err != nil {
		return err
	}

	aVal, ok := a.(IntValue)
	if !ok {
		return errors.New("unexpected value on stack")
	}

	b, err := ec.Stack.pop()
	if err != nil {
		return err
	}

	bVal, ok := b.(IntValue)
	if !ok {
		return errors.New("unexpected value on stack")
	}

	ec.Stack.Push(IntValue{val: aVal.val + bVal.val})
	return nil
}

type Mul struct {
}

type Push1 struct{}

func (p Push1) Execute(ec *EvmInterpreter) error {
	val := ec.next()
	parsed, err := strconv.ParseInt(val, 16, 8)
	if err != nil {
		return err
	}

	intVal := int(parsed)

	ec.Stack.Push(IntValue{val: intVal})
	return nil
}

type Swap1 struct {
}

func (swap Swap1) Execute(ec *EvmInterpreter) error {
	if ec.Stack.Size() < 2 {
		return errors.New("not enough elements on the stack")
	}

	a, err := ec.Stack.pop()
	if err != nil {
		return err
	}

	b, err := ec.Stack.pop()
	if err != nil {
		return err
	}

	ec.Stack.Push(a)
	ec.Stack.Push(b)
	return nil
}
