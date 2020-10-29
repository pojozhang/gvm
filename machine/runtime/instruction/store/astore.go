package store

import (
	"gvm/machine/runtime"
	"gvm/machine/runtime/instruction"
)

func storeRef(frame *runtime.Frame, index uint) {
	frame.GetLocalVars().SetRef(index, frame.GetOperandStack().PopRef())
}

type AStore struct {
	instruction.Index8Instruction
}

func (i *AStore) Execute(frame *runtime.Frame) {
	storeRef(frame, i.GetIndex())
}

type AStore0 struct {
	instruction.NoOperandsInstruction
}

func (i *AStore0) Execute(frame *runtime.Frame) {
	storeRef(frame, 0)
}

type AStore1 struct {
	instruction.NoOperandsInstruction
}

func (i *AStore1) Execute(frame *runtime.Frame) {
	storeRef(frame, 1)
}

type AStore2 struct {
	instruction.NoOperandsInstruction
}

func (i *AStore2) Execute(frame *runtime.Frame) {
	storeRef(frame, 2)
}

type AStore3 struct {
	instruction.NoOperandsInstruction
}

func (i *AStore3) Execute(frame *runtime.Frame) {
	storeRef(frame, 3)
}

type WAStore struct {
	instruction.Index16Instruction
}

func (i *WAStore) Execute(frame *runtime.Frame) {
	storeRef(frame, i.GetIndex())
}
