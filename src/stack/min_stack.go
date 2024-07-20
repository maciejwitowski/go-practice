package stack

/*
Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

Implement the MinStack class:

MinStack() initializes the stack object.
void push(int val) pushes the element val onto the stack.
void pop() removes the element on the top of the stack.
int top() gets the top element of the stack.
int getMin() retrieves the minimum element in the stack.
You must implement a solution with O(1) time complexity for each function.

*/

type MinStack struct {
	items    *internalStack
	minItems *internalStack
}

func Constructor() MinStack {
	return MinStack{
		items:    &internalStack{make([]int, 0)},
		minItems: &internalStack{make([]int, 0)},
	}
}

func (this *MinStack) Push(val int) {
	this.items.push(val)

	if this.minItems.isEmpty() || val <= this.minItems.peek() {
		this.minItems.push(val)
	}
}

func (this *MinStack) Pop() {
	if this.items.peek() == this.minItems.peek() {
		this.minItems.pop()
	}

	this.items.pop()
}

func (this *MinStack) Top() int {
	return this.items.peek()
}

func (this *MinStack) GetMin() int {
	return this.minItems.peek()
}

type internalStack struct {
	items []int
}

func (s *internalStack) push(val int) {
	s.items = append(s.items, val)
}

func (s *internalStack) pop() {
	if len(s.items) == 0 {
		return
	}
	s.items = s.items[:len(s.items)-1]
}

func (s *internalStack) peek() int {
	return s.items[len(s.items)-1]
}

func (s *internalStack) isEmpty() bool {
	return len(s.items) == 0
}
