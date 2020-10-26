package comparison

import (
	"gvm/machine/runtime"
	instruction "gvm/machine/runtime/instructions"
)

type FCmpG struct {
	instruction.NoOperandsInstruction
}

func (i *FCmpG) Execute(frame *runtime.Frame) {
	compareFloat(frame, true)
}

type FCmpL struct {
	instruction.NoOperandsInstruction
}

func (i *FCmpL) Execute(frame *runtime.Frame) {
	compareFloat(frame, false)
}

func compareFloat(frame *runtime.Frame, gFlag bool) {
	stack := frame.GetOperandStack()
	v1 := stack.PopFloat()
	v2 := stack.PopFloat()
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
