package conversion

import (
	"gvm/machine/runtime"
	instruction "gvm/machine/runtime/instruction"
)

type L2I struct {
	instruction.NoOperandsInstruction
}

func (i *L2I) Execute(frame *runtime.Frame) {
	stack := frame.GetOperandStack()
	stack.PushInt(int32(stack.PopLong()))
}

type L2F struct {
	instruction.NoOperandsInstruction
}

func (i *L2F) Execute(frame *runtime.Frame) {
	stack := frame.GetOperandStack()
	stack.PushFloat(float32(stack.PopLong()))
}

type L2D struct {
	instruction.NoOperandsInstruction
}

func (i *L2D) Execute(frame *runtime.Frame) {
	stack := frame.GetOperandStack()
	stack.PushDouble(float64(stack.PopLong()))
}
