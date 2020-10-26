package comparison

import (
	"gvm/machine/runtime"
	instruction "gvm/machine/runtime/instruction"
)

type IfACmpEq struct {
	instruction.BranchInstruction
}

func (i *IfACmpEq) Execute(frame *runtime.Frame) {
	stack := frame.GetOperandStack()
	if stack.PopRef() == stack.PopRef() {
		i.Branch(frame, i.GetOffset())
	}
}

type IfACmpNe struct {
	instruction.BranchInstruction
}

func (i *IfACmpNe) Execute(frame *runtime.Frame) {
	stack := frame.GetOperandStack()
	if stack.PopRef() != stack.PopRef() {
		i.Branch(frame, i.GetOffset())
	}
}
