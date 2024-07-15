package sedgewick

import (
	"fmt"
	"slices"
	"unicode"
)

var supportedOperators = []rune{'+', '-', '*'}

func Evaluate(expr string) (int, error) {
	operands := NewStack[int](0)
	operators := NewStack[rune](0)

	for i := 0; i < len(expr); i++ {
		s := rune(expr[i])
		if unicode.IsDigit(s) {
			num := 0
			for i < len(expr) && unicode.IsDigit(rune(expr[i])) {
				num = num*10 + int(expr[i]-'0')
				i++
			}
			i--
			operands.Push(num)
		} else if slices.Contains(supportedOperators, s) {
			for operators.Len() > 0 && precedence(operators.Peek()) >= precedence(s) {
				if err := applyOp(operands, operators); err != nil {
					return 0, err
				}
			}
			operators.Push(s)
		} else if s == '(' {
			operators.Push(s)
		} else if s == ')' {
			for operators.Len() > 0 && operators.Peek() != '(' {
				if err := applyOp(operands, operators); err != nil {
					return 0, err
				}
			}
			if operators.Len() == 0 {
				return 0, fmt.Errorf("mismatched parentheses")
			}
			operators.Pop() // Remove '('
		}
	}

	for operators.Len() > 0 {
		if err := applyOp(operands, operators); err != nil {
			return 0, err
		}
	}

	if operands.Len() != 1 {
		return 0, fmt.Errorf("invalid expression")
	}

	return operands.Pop()
}

func precedence(op rune) int {
	switch op {
	case '+', '-':
		return 1
	case '*', '/':
		return 2
	}
	return 0
}

func applyOp(operands *Stack[int], operators *Stack[rune]) error {
	if operands.Len() < 2 {
		return fmt.Errorf("not enough operands")
	}
	op, err := operators.Pop()
	if err != nil {
		return err
	}
	b, err := operands.Pop()
	if err != nil {
		return err
	}
	a, err := operands.Pop()
	if err != nil {
		return err
	}
	switch op {
	case '+':
		operands.Push(a + b)
	case '-':
		operands.Push(a - b)
	case '*':
		operands.Push(a * b)
	case '/':
		if b == 0 {
			return fmt.Errorf("division by zero")
		}
		operands.Push(a / b)
	default:
		return fmt.Errorf("unsupported operator %c", op)
	}
	return nil
}

type ItemType interface {
	int | float64 | rune
}
type Stack[T ItemType] struct {
	items []T
}

func NewStack[T ItemType](initialCapacity int) *Stack[T] {
	return &Stack[T]{
		items: make([]T, initialCapacity),
	}
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, error) {
	lastIndex := len(s.items) - 1
	last := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return last, nil
}

func (s *Stack[T]) Peek() T {
	return s.items[len(s.items)-1]
}

func (s *Stack[T]) isEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack[T]) Len() int {
	return len(s.items)
}
