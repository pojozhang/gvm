package conversion

import (
	"gvm/machine/runtime"
	instruction "gvm/machine/runtime/instructions"
)

type F2I struct {
	instruction.NoOperandsInstruction
}

func (i *F2I) Execute(frame *runtime.Frame) {
	stack := frame.GetOperandStack()
	stack.PushInt(int32(stack.PopFloat()))
}

type F2D struct {
	instruction.NoOperandsInstruction
}

func (i *F2D) Execute(frame *runtime.Frame) {
	stack := frame.GetOperandStack()
	stack.PushDouble(float64(stack.PopFloat()))
}

type F2L struct {
	instruction.NoOperandsInstruction
}

func (i *F2L) Execute(frame *runtime.Frame) {
	stack := frame.GetOperandStack()
	stack.PushLong(int64(stack.PopFloat()))
}
