package stack

import (
	"fmt"
	"strconv"
)

/*
You are given an array of strings tokens that represents an arithmetic expression in a Reverse Polish Notation.

Evaluate the expression. Return an integer that represents the value of the expression.

Note that:

The valid operators are '+', '-', '*', and '/'.
Each operand may be an integer or another expression.
The division between two integers always truncates toward zero.
There will not be any division by zero.
The input represents a valid arithmetic expression in a reverse polish notation.
The answer and all the intermediate calculations can be represented in a 32-bit integer.

*/

func evalRPN(tokens []string) int {
	stack := genericStack[Token]{}

	for _, t := range tokens {
		token, err := parse(t)
		if err != nil {
			return -1
		}

		token.Execute(&stack)
	}

	result, stackNotEmpty := stack.pop()
	if !stackNotEmpty {
		return -1
	}

	number, ok := result.(Number)
	if !ok {
		return -1
	}

	return int(number)
}

type Token interface {
	Execute(stack *genericStack[Token])
}

type Number int

func (n Number) Execute(stack *genericStack[Token]) {
	stack.push(n)
}

type Add struct{}

func (n Add) Execute(stack *genericStack[Token]) {
	twoNumbersOperation(stack, func(a Number, b Number) Number {
		return a + b
	})
}

type Subtract struct{}

func (n Subtract) Execute(stack *genericStack[Token]) {
	twoNumbersOperation(stack, func(a Number, b Number) Number {
		return a - b
	})
}

type Divide struct{}

func (n Divide) Execute(stack *genericStack[Token]) {
	twoNumbersOperation(stack, func(a Number, b Number) Number {
		return a / b
	})
}

type Multiply struct{}

func (n Multiply) Execute(stack *genericStack[Token]) {
	twoNumbersOperation(stack, func(a Number, b Number) Number {
		return a * b
	})
}

func twoNumbersOperation(stack *genericStack[Token], op func(a Number, b Number) Number) {
	b, ok2 := stack.pop()
	a, ok1 := stack.pop()
	if !ok1 || !ok2 {
		fmt.Println("Error: Not enough elements on the stack to perform addition")
		return
	}

	aNum, ok1 := a.(Number)
	bNum, ok2 := b.(Number)

	if !ok1 || !ok2 {
		fmt.Println("Error: Elements on the stack are not of type Number")
		return
	}

	stack.push(op(aNum, bNum))
}

func parse(s string) (Token, error) {
	switch s {
	case "+":
		return Add{}, nil
	case "-":
		return Subtract{}, nil
	case "/":
		return Divide{}, nil
	case "*":
		return Multiply{}, nil
	default:
		num, err := strconv.Atoi(s)

		if err == nil {
			return Number(num), nil
		} else {
			return nil, err
		}
	}
}

type genericStack[T any] struct {
	items []T
}

func (s *genericStack[T]) push(t T) {
	s.items = append(s.items, t)
}

func (s *genericStack[T]) pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}

	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, true
}

func (s *genericStack[T]) peek() T {
	return s.items[len(s.items)-1]
}

func (s *genericStack[T]) isEmpty() bool {
	return len(s.items) == 0
}
