package load

import (
	"gvm/machine/runtime"
	"gvm/machine/runtime/instruction"
)

func loadInt(frame *runtime.Frame, index uint) {
	frame.GetOperandStack().PushInt(frame.GetLocalVars().GetInt(index))
}

type ILoad struct {
	instruction.Index8Instruction
}

func (i *ILoad) Execute(frame *runtime.Frame) {
	loadInt(frame, i.GetIndex())
}

type ILoad0 struct {
	instruction.NoOperandsInstruction
}

func (i *ILoad0) Execute(frame *runtime.Frame) {
	loadInt(frame, 0)
}

type ILoad1 struct {
	instruction.NoOperandsInstruction
}

func (i *ILoad1) Execute(frame *runtime.Frame) {
	loadInt(frame, 1)
}

type ILoad2 struct {
	instruction.NoOperandsInstruction
}

func (i *ILoad2) Execute(frame *runtime.Frame) {
	loadInt(frame, 2)
}

type ILoad3 struct {
	instruction.NoOperandsInstruction
}

func (i *ILoad3) Execute(frame *runtime.Frame) {
	loadInt(frame, 3)
}

type WILoad struct {
	instruction.Index16Instruction
}

func (i *WILoad) Execute(frame *runtime.Frame) {
	loadInt(frame, i.GetIndex())
}
