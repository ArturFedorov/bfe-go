package min_stack

type MinStack struct {
	stack []int
	min   []int
	top   int
}

func Constructor() MinStack {
	return MinStack{[]int{}, []int{}, -1}
}
func (s *MinStack) Push(val int) {
	if s.top == -1 {
		s.min = append(s.min, val)
	} else {
		s.min = append(s.min, minVal(val, s.min[s.top]))
	}

	s.stack = append(s.stack, val)
	s.top++
}

func (s *MinStack) Pop() {
	s.stack = s.stack[:s.top]
	s.min = s.min[:s.top]
	s.top--
}
func (s *MinStack) Top() int {
	return s.stack[s.top]
}

func (s *MinStack) GetMin() int {
	return s.min[s.top]
}

func minVal(a, b int) int {
	if a < b {
		return a
	}

	return b
}
