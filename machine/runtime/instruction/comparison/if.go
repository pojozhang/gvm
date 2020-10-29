package comparison

import (
	"gvm/machine/runtime"
	instruction "gvm/machine/runtime/instruction"
)

type IfEq struct {
	instruction.JumpInstruction
}

func (i *IfEq) Execute(frame *runtime.Frame) {
	if frame.GetOperandStack().PopInt() == 0 {
		frame.Jump(i.GetOffset())
	}
}

type IfNe struct {
	instruction.JumpInstruction
}

func (i *IfNe) Execute(frame *runtime.Frame) {
	if frame.GetOperandStack().PopInt() != 0 {
		frame.Jump(i.GetOffset())
	}
}

type IfLt struct {
	instruction.JumpInstruction
}

func (i *IfLt) Execute(frame *runtime.Frame) {
	if frame.GetOperandStack().PopInt() < 0 {
		frame.Jump(i.GetOffset())
	}
}

type IfLe struct {
	instruction.JumpInstruction
}

func (i *IfLe) Execute(frame *runtime.Frame) {
	if frame.GetOperandStack().PopInt() <= 0 {
		frame.Jump(i.GetOffset())
	}
}

type IfGt struct {
	instruction.JumpInstruction
}

func (i *IfGt) Execute(frame *runtime.Frame) {
	if frame.GetOperandStack().PopInt() > 0 {
		frame.Jump(i.GetOffset())
	}
}

type IfGe struct {
	instruction.JumpInstruction
}

func (i *IfGe) Execute(frame *runtime.Frame) {
	if frame.GetOperandStack().PopInt() >= 0 {
		frame.Jump(i.GetOffset())
	}
}
