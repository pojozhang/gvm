package load

import (
	"gvm/machine/runtime"
	"gvm/machine/runtime/instruction"
)

func loadDouble(frame *runtime.Frame, index uint) {
	frame.GetOperandStack().PushDouble(frame.GetLocalVars().GetDouble(index))
}

type DLoad struct {
	instruction.Index8Instruction
}

func (i *DLoad) Execute(frame *runtime.Frame) {
	loadDouble(frame, i.GetIndex())
}

type DLoad0 struct {
	instruction.NoOperandsInstruction
}

func (i *DLoad0) Execute(frame *runtime.Frame) {
	loadDouble(frame, 0)
}

type DLoad1 struct {
	instruction.NoOperandsInstruction
}

func (i *DLoad1) Execute(frame *runtime.Frame) {
	loadDouble(frame, 1)
}

type DLoad2 struct {
	instruction.NoOperandsInstruction
}

func (i *DLoad2) Execute(frame *runtime.Frame) {
	loadDouble(frame, 2)
}

type DLoad3 struct {
	instruction.NoOperandsInstruction
}

func (i *DLoad3) Execute(frame *runtime.Frame) {
	loadDouble(frame, 3)
}

type WDLoad struct {
	instruction.Index16Instruction
}

func (i *WDLoad) Execute(frame *runtime.Frame) {
	loadDouble(frame, i.GetIndex())
}
