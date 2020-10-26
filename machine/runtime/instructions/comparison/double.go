package comparison

import (
	"gvm/machine/runtime"
	instruction "gvm/machine/runtime/instructions"
)

type DCmpG struct {
	instruction.NoOperandsInstruction
}

func (i *DCmpG) Execute(frame *runtime.Frame) {
	compareDouble(frame, true)
}

type DCmpL struct {
	instruction.NoOperandsInstruction
}

func (i *DCmpL) Execute(frame *runtime.Frame) {
	compareDouble(frame, false)
}

func compareDouble(frame *runtime.Frame, gFlag bool) {
	stack := frame.GetOperandStack()
	v1 := stack.PopDouble()
	v2 := stack.PopDouble()
	if v2 > v1 {
		stack.PushInt(1)
	} else if v1 < v2 {
		stack.PushInt(-1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else if gFlag {
		stack.PushInt(1)
	} else {
		stack.PushInt(-1)
	}
}
