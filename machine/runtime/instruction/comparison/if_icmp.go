package comparison

import (
	"gvm/machine/runtime"
	instruction "gvm/machine/runtime/instruction"
)

type IfICmpEq struct {
	instruction.JumpInstruction
}

func (i *IfICmpEq) Execute(frame *runtime.Frame) {
	stack := frame.GetOperandStack()
	if stack.PopInt() == stack.PopInt() {
		frame.Jump(i.GetOffset())
	}
}

type IfICmpNe struct {
	instruction.JumpInstruction
}

func (i *IfICmpNe) Execute(frame *runtime.Frame) {
	stack := frame.GetOperandStack()
	if stack.PopInt() != stack.PopInt() {
		frame.Jump(i.GetOffset())
	}
}

type IfICmpLt struct {
	instruction.JumpInstruction
}

func (i *IfICmpLt) Execute(frame *runtime.Frame) {
	stack := frame.GetOperandStack()
	v1, v2 := stack.PopInt(), stack.PopInt()
	if v2 < v1 {
		frame.Jump(i.GetOffset())
	}
}

type IfICmpLe struct {
	instruction.JumpInstruction
}

func (i *IfICmpLe) Execute(frame *runtime.Frame) {
	stack := frame.GetOperandStack()
	v1, v2 := stack.PopInt(), stack.PopInt()
	if v2 <= v1 {
		frame.Jump(i.GetOffset())
	}
}

type IfICmpGt struct {
	instruction.JumpInstruction
}

func (i *IfICmpGt) Execute(frame *runtime.Frame) {
	stack := frame.GetOperandStack()
	v1, v2 := stack.PopInt(), stack.PopInt()
	if v2 > v1 {
		frame.Jump(i.GetOffset())
	}
}

type IfICmpGe struct {
	instruction.JumpInstruction
}

func (i *IfICmpGe) Execute(frame *runtime.Frame) {
	stack := frame.GetOperandStack()
	v1, v2 := stack.PopInt(), stack.PopInt()
	if v2 >= v1 {
		frame.Jump(i.GetOffset())
	}
}
