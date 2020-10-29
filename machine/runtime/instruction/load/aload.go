package load

import (
	"gvm/machine/runtime"
	"gvm/machine/runtime/instruction"
)

func loadRef(frame *runtime.Frame, index uint) {
	frame.GetOperandStack().PushRef(frame.GetLocalVars().GetRef(index))
}

type ALoad struct {
	instruction.Index8Instruction
}

func (i *ALoad) Execute(frame *runtime.Frame) {
	loadRef(frame, i.GetIndex())
}

type ALoad0 struct {
	instruction.NoOperandsInstruction
}

func (i *ALoad0) Execute(frame *runtime.Frame) {
	loadRef(frame, 0)
}

type ALoad1 struct {
	instruction.NoOperandsInstruction
}

func (i *ALoad1) Execute(frame *runtime.Frame) {
	loadRef(frame, 1)
}

type ALoad2 struct {
	instruction.NoOperandsInstruction
}

func (i *ALoad2) Execute(frame *runtime.Frame) {
	loadRef(frame, 2)
}

type ALoad3 struct {
	instruction.NoOperandsInstruction
}

func (i *ALoad3) Execute(frame *runtime.Frame) {
	loadRef(frame, 3)
}

type WALoad struct {
	instruction.Index16Instruction
}

func (i *WALoad) Execute(frame *runtime.Frame) {
	loadRef(frame, i.GetIndex())
}
