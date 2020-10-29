package store

import (
	"gvm/machine/runtime"
	"gvm/machine/runtime/instruction"
)

func storeFloat(frame *runtime.Frame, index uint) {
	frame.GetLocalVars().SetFloat(index, frame.GetOperandStack().PopFloat())
}

type FStore struct {
	instruction.Index8Instruction
}

func (i *FStore) Execute(frame *runtime.Frame) {
	storeFloat(frame, i.GetIndex())
}

type FStore0 struct {
	instruction.NoOperandsInstruction
}

func (i *FStore0) Execute(frame *runtime.Frame) {
	storeFloat(frame, 0)
}

type FStore1 struct {
	instruction.NoOperandsInstruction
}

func (i *FStore1) Execute(frame *runtime.Frame) {
	storeFloat(frame, 1)
}

type FStore2 struct {
	instruction.NoOperandsInstruction
}

func (i *FStore2) Execute(frame *runtime.Frame) {
	storeFloat(frame, 2)
}

type FStore3 struct {
	instruction.NoOperandsInstruction
}

func (i *FStore3) Execute(frame *runtime.Frame) {
	storeFloat(frame, 3)
}

type WFStore struct {
	instruction.Index16Instruction
}

func (i *WFStore) Execute(frame *runtime.Frame) {
	storeFloat(frame, i.GetIndex())
}
