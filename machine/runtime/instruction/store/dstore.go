package store

import (
	"gvm/machine/runtime"
	"gvm/machine/runtime/instruction"
)

func storeDouble(frame *runtime.Frame, index uint) {
	frame.GetLocalVars().SetDouble(index, frame.GetOperandStack().PopDouble())
}

type DStore struct {
	instruction.Index8Instruction
}

func (i *DStore) Execute(frame *runtime.Frame) {
	storeDouble(frame, i.GetIndex())
}

type DStore0 struct {
	instruction.NoOperandsInstruction
}

func (i *DStore0) Execute(frame *runtime.Frame) {
	storeDouble(frame, 0)
}

type DStore1 struct {
	instruction.NoOperandsInstruction
}

func (i *DStore1) Execute(frame *runtime.Frame) {
	storeDouble(frame, 1)
}

type DStore2 struct {
	instruction.NoOperandsInstruction
}

func (i *DStore2) Execute(frame *runtime.Frame) {
	storeDouble(frame, 2)
}

type DStore3 struct {
	instruction.NoOperandsInstruction
}

func (i *DStore3) Execute(frame *runtime.Frame) {
	storeDouble(frame, 3)
}

type WDStore struct {
	instruction.Index16Instruction
}

func (i *WDStore) Execute(frame *runtime.Frame) {
	storeDouble(frame, i.GetIndex())
}
