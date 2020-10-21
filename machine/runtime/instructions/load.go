package instruction

import "gvm/machine/runtime"

func loadInt(frame *runtime.Frame, index uint) {
	frame.GetOperandStack().PushInt(frame.GetLocalVars().GetInt(index))
}

type ILoad struct {
	Index8Instruction
}

func (i *ILoad) Execute(frame *runtime.Frame) {
	loadInt(frame, i.index)
}

type ILoad1 struct {
	NoOperandsInstruction
}

func (i *ILoad1) Execute(frame *runtime.Frame) {
	loadInt(frame, 1)
}

type ILoad2 struct {
	NoOperandsInstruction
}

func (i *ILoad2) Execute(frame *runtime.Frame) {
	loadInt(frame, 2)
}

type ILoad3 struct {
	NoOperandsInstruction
}

func (i *ILoad3) Execute(frame *runtime.Frame) {
	loadInt(frame, 3)
}
