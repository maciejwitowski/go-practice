package sedgewick

import (
	"testing"
)

func TestReverse(t *testing.T) {
	expr := "(1+((2+3)*(4*5)))"
	result, _ := Evaluate(expr)

	if result != 101 {
		t.Errorf("incorrect result %d", result)
	}
}
