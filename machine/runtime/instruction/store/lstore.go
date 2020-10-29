package store

import (
	"gvm/machine/runtime"
	"gvm/machine/runtime/instruction"
)

func storeLong(frame *runtime.Frame, index uint) {
	frame.GetLocalVars().SetLong(index, frame.GetOperandStack().PopLong())
}

type LStore struct {
	instruction.Index8Instruction
}

func (i *LStore) Execute(frame *runtime.Frame) {
	storeLong(frame, i.GetIndex())
}

type LStore0 struct {
	instruction.NoOperandsInstruction
}

func (i *LStore0) Execute(frame *runtime.Frame) {
	storeLong(frame, 0)
}

type LStore1 struct {
	instruction.NoOperandsInstruction
}

func (i *LStore1) Execute(frame *runtime.Frame) {
	storeLong(frame, 1)
}

type LStore2 struct {
	instruction.NoOperandsInstruction
}

func (i *LStore2) Execute(frame *runtime.Frame) {
	storeLong(frame, 2)
}

type LStore3 struct {
	instruction.NoOperandsInstruction
}

func (i *LStore3) Execute(frame *runtime.Frame) {
	storeLong(frame, 3)
}

type WLStore struct {
	instruction.Index16Instruction
}

func (i *WLStore) Execute(frame *runtime.Frame) {
	storeLong(frame, i.GetIndex())
}
