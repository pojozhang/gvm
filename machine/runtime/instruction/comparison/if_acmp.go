package comparison

import (
	"gvm/machine/runtime"
	instruction "gvm/machine/runtime/instruction"
)

type IfACmpEq struct {
	instruction.JumpInstruction
}

func (i *IfACmpEq) Execute(frame *runtime.Frame) {
	stack := frame.GetOperandStack()
	if stack.PopRef() == stack.PopRef() {
		frame.Jump(i.GetOffset())
	}
}

type IfACmpNe struct {
	instruction.JumpInstruction
}

func (i *IfACmpNe) Execute(frame *runtime.Frame) {
	stack := frame.GetOperandStack()
	if stack.PopRef() != stack.PopRef() {
		frame.Jump(i.GetOffset())
	}
}
