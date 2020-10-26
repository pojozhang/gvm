package math

import (
	"gvm/machine/runtime"
	instruction "gvm/machine/runtime/instruction"
)

type IAnd struct {
	instruction.NoOperandsInstruction
}

func (i *IAnd) Execute(frame *runtime.Frame) {
	stack := frame.GetOperandStack()
	v1, v2 := stack.PopInt(), stack.PopInt()
	stack.PushInt(v2 & v1)
}

type LAnd struct {
	instruction.NoOperandsInstruction
}

func (i *LAnd) Execute(frame *runtime.Frame) {
	stack := frame.GetOperandStack()
	v1, v2 := stack.PopLong(), stack.PopLong()
	stack.PushLong(v2 & v1)
}
