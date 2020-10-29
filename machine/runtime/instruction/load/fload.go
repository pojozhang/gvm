package load

import (
	"gvm/machine/runtime"
	"gvm/machine/runtime/instruction"
)

func loadFloat(frame *runtime.Frame, index uint) {
	frame.GetOperandStack().PushFloat(frame.GetLocalVars().GetFloat(index))
}

type FLoad struct {
	instruction.Index8Instruction
}

func (i *FLoad) Execute(frame *runtime.Frame) {
	loadFloat(frame, i.GetIndex())
}

type FLoad0 struct {
	instruction.NoOperandsInstruction
}

func (i *FLoad0) Execute(frame *runtime.Frame) {
	loadFloat(frame, 0)
}

type FLoad1 struct {
	instruction.NoOperandsInstruction
}

func (i *FLoad1) Execute(frame *runtime.Frame) {
	loadFloat(frame, 1)
}

type FLoad2 struct {
	instruction.NoOperandsInstruction
}

func (i *FLoad2) Execute(frame *runtime.Frame) {
	loadFloat(frame, 2)
}

type FLoad3 struct {
	instruction.NoOperandsInstruction
}

func (i *FLoad3) Execute(frame *runtime.Frame) {
	loadFloat(frame, 3)
}

type WFLoad struct {
	instruction.Index16Instruction
}

func (i *WFLoad) Execute(frame *runtime.Frame) {
	loadFloat(frame, i.GetIndex())
}
