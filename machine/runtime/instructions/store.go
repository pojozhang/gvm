package instruction

import "gvm/machine/runtime"

func storeLong(frame *runtime.Frame, index uint) {
	frame.GetLocalVars().SetLong(index, frame.GetOperandStack().PopLong())
}

type LStore struct {
	Index8Instruction
}

func (i *LStore) Execute(frame *runtime.Frame) {
	storeLong(frame, i.index)
}

type LStore0 struct {
	NoOperandsInstruction
}

func (i *LStore0) Execute(frame *runtime.Frame) {
	storeLong(frame, 0)
}

type LStore1 struct {
	NoOperandsInstruction
}

func (i *LStore1) Execute(frame *runtime.Frame) {
	storeLong(frame, 1)
}

type LStore2 struct {
	NoOperandsInstruction
}

func (i *LStore2) Execute(frame *runtime.Frame) {
	storeLong(frame, 2)
}

type LStore3 struct {
	NoOperandsInstruction
}

func (i *LStore3) Execute(frame *runtime.Frame) {
	storeLong(frame, 3)
}
