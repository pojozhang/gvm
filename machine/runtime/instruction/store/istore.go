package store

import (
	"gvm/machine/runtime"
	"gvm/machine/runtime/instruction"
)

func storeInt(frame *runtime.Frame, index uint) {
	frame.GetLocalVars().SetInt(index, frame.GetOperandStack().PopInt())
}

type IStore struct {
	instruction.Index8Instruction
}

func (i *IStore) Execute(frame *runtime.Frame) {
	storeInt(frame, i.GetIndex())
}

type IStore0 struct {
	instruction.NoOperandsInstruction
}

func (i *IStore0) Execute(frame *runtime.Frame) {
	storeInt(frame, 0)
}

type IStore1 struct {
	instruction.NoOperandsInstruction
}

func (i *IStore1) Execute(frame *runtime.Frame) {
	storeInt(frame, 1)
}

type IStore2 struct {
	instruction.NoOperandsInstruction
}

func (i *IStore2) Execute(frame *runtime.Frame) {
	storeInt(frame, 2)
}

type IStore3 struct {
	instruction.NoOperandsInstruction
}

func (i *IStore3) Execute(frame *runtime.Frame) {
	storeInt(frame, 3)
}

type WIStore struct {
	instruction.Index16Instruction
}

func (i *WIStore) Execute(frame *runtime.Frame) {
	storeInt(frame, i.GetIndex())
}
