package comparison

import (
	"gvm/machine/runtime"
	instruction "gvm/machine/runtime/instructions"
)

type IfICmpEq struct {
	instruction.BranchInstruction
}

func (i *IfICmpEq) Execute(frame *runtime.Frame) {
	stack := frame.GetOperandStack()
	if stack.PopInt() == stack.PopInt() {
		i.Branch(frame, i.GetOffset())
	}
}

type IfICmpNe struct {
	instruction.BranchInstruction
}

func (i *IfICmpNe) Execute(frame *runtime.Frame) {
	stack := frame.GetOperandStack()
	if stack.PopInt() != stack.PopInt() {
		i.Branch(frame, i.GetOffset())
	}
}

type IfICmpLt struct {
	instruction.BranchInstruction
}

func (i *IfICmpLt) Execute(frame *runtime.Frame) {
	stack := frame.GetOperandStack()
	v1, v2 := stack.PopInt(), stack.PopInt()
	if v2 < v1 {
		i.Branch(frame, i.GetOffset())
	}
}

type IfICmpLe struct {
	instruction.BranchInstruction
}

func (i *IfICmpLe) Execute(frame *runtime.Frame) {
	stack := frame.GetOperandStack()
	v1, v2 := stack.PopInt(), stack.PopInt()
	if v2 <= v1 {
		i.Branch(frame, i.GetOffset())
	}
}

type IfICmpGt struct {
	instruction.BranchInstruction
}

func (i *IfICmpGt) Execute(frame *runtime.Frame) {
	stack := frame.GetOperandStack()
	v1, v2 := stack.PopInt(), stack.PopInt()
	if v2 > v1 {
		i.Branch(frame, i.GetOffset())
	}
}

type IfICmpGe struct {
	instruction.BranchInstruction
}

func (i *IfICmpGe) Execute(frame *runtime.Frame) {
	stack := frame.GetOperandStack()
	v1, v2 := stack.PopInt(), stack.PopInt()
	if v2 >= v1 {
		i.Branch(frame, i.GetOffset())
	}
}
