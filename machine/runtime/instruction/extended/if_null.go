package extended

import (
	"gvm/machine/runtime"
	"gvm/machine/runtime/instruction"
)

type IfNull struct {
	instruction.JumpInstruction
}

func (i *IfNull) Execute(frame *runtime.Frame) {
	if frame.GetOperandStack().PopRef() == nil {
		frame.Jump(i.GetOffset())
	}
}

type IfNonNull struct {
	instruction.JumpInstruction
}

func (i *IfNonNull) Execute(frame *runtime.Frame) {
	if frame.GetOperandStack().PopRef() != nil {
		frame.Jump(i.GetOffset())
	}
}
