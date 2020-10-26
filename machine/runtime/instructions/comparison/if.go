package comparison

import (
	"gvm/machine/runtime"
	instruction "gvm/machine/runtime/instructions"
)

type IfEq struct {
	instruction.BranchInstruction
}

func (i *IfEq) Execute(frame *runtime.Frame) {
	if frame.GetOperandStack().PopInt() == 0 {
		i.Branch(frame, i.GetOffset())
	}
}

type IfNe struct {
	instruction.BranchInstruction
}

func (i *IfNe) Execute(frame *runtime.Frame) {
	if frame.GetOperandStack().PopInt() != 0 {
		i.Branch(frame, i.GetOffset())
	}
}

type IfLt struct {
	instruction.BranchInstruction
}

func (i *IfLt) Execute(frame *runtime.Frame) {
	if frame.GetOperandStack().PopInt() < 0 {
		i.Branch(frame, i.GetOffset())
	}
}

type IfLe struct {
	instruction.BranchInstruction
}

func (i *IfLe) Execute(frame *runtime.Frame) {
	if frame.GetOperandStack().PopInt() <= 0 {
		i.Branch(frame, i.GetOffset())
	}
}

type IfGt struct {
	instruction.BranchInstruction
}

func (i *IfGt) Execute(frame *runtime.Frame) {
	if frame.GetOperandStack().PopInt() > 0 {
		i.Branch(frame, i.GetOffset())
	}
}

type IfGe struct {
	instruction.BranchInstruction
}

func (i *IfGe) Execute(frame *runtime.Frame) {
	if frame.GetOperandStack().PopInt() >= 0 {
		i.Branch(frame, i.GetOffset())
	}
}
