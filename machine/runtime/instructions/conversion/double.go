package conversion

import (
	"gvm/machine/runtime"
	instruction "gvm/machine/runtime/instructions"
)

type D2I struct {
	instruction.NoOperandsInstruction
}

func (i *D2I) Execute(frame *runtime.Frame) {
	stack := frame.GetOperandStack()
	stack.PushInt(int32(stack.PopDouble()))
}

type D2F struct {
	instruction.NoOperandsInstruction
}

func (i *D2F) Execute(frame *runtime.Frame) {
	stack := frame.GetOperandStack()
	stack.PushFloat(float32(stack.PopDouble()))
}

type D2L struct {
	instruction.NoOperandsInstruction
}

func (i *D2L) Execute(frame *runtime.Frame) {
	stack := frame.GetOperandStack()
	stack.PushLong(int64(stack.PopDouble()))
}
