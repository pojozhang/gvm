package instruction

import "gvm/machine/runtime"

// 复制栈顶单个变量。
type Dup struct {
	NoOperandsInstruction
}

func (i *Dup) Execute(frame *runtime.Frame) {
	stack := frame.GetOperandStack()
	slot := stack.PopSlot()
	stack.PushSlot(slot)
	stack.PushSlot(slot)
}
