package runtime

type Thread struct {
	pc    int
	stack *Stack
}

func (t *Thread) GetPC() int {
	return t.pc
}

func NewThread() *Thread {
	return &Thread{
		stack: newStack(1024),
	}
}
