package load

import (
	"gvm/machine/runtime"
	"gvm/machine/runtime/instruction"
)

func loadLong(frame *runtime.Frame, index uint) {
	frame.GetOperandStack().PushLong(frame.GetLocalVars().GetLong(index))
}

type LLoad struct {
	instruction.Index8Instruction
}

func (i *LLoad) Execute(frame *runtime.Frame) {
	loadLong(frame, i.GetIndex())
}

type LLoad0 struct {
	instruction.NoOperandsInstruction
}

func (i *LLoad0) Execute(frame *runtime.Frame) {
	loadLong(frame, 0)
}

type LLoad1 struct {
	instruction.NoOperandsInstruction
}

func (i *LLoad1) Execute(frame *runtime.Frame) {
	loadLong(frame, 1)
}

type LLoad2 struct {
	instruction.NoOperandsInstruction
}

func (i *LLoad2) Execute(frame *runtime.Frame) {
	loadLong(frame, 2)
}

type LLoad3 struct {
	instruction.NoOperandsInstruction
}

func (i *LLoad3) Execute(frame *runtime.Frame) {
	loadLong(frame, 3)
}

type WLLoad struct {
	instruction.Index16Instruction
}

func (i *WLLoad) Execute(frame *runtime.Frame) {
	loadLong(frame, i.GetIndex())
}
