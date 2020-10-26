package comparison

import (
	"gvm/machine/runtime"
	instruction "gvm/machine/runtime/instruction"
)

type LCmp struct {
	instruction.NoOperandsInstruction
}

func (i *LCmp) Execute(frame *runtime.Frame) {
	stack := frame.GetOperandStack()
	v1 := stack.PopLong()
	v2 := stack.PopLong()
	if v2 > v1 {
		stack.PushInt(1)
	} else if v1 < v2 {
		stack.PushInt(-1)
	} else {
		stack.PushInt(0)
	}
}
