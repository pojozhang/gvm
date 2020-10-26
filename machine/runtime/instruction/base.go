package instruction

import "gvm/machine/runtime"

type NoOperandsInstruction struct {
}

func (i *NoOperandsInstruction) FetchOperands(reader *runtime.BytecodeReader) {
}

func (i *NoOperandsInstruction) Execute(frame *runtime.Frame) {
}

type BranchInstruction struct {
	offset int
}

func (i *BranchInstruction) GetOffset() int {
	return i.offset
}

func (i *BranchInstruction) FetchOperands(reader *runtime.BytecodeReader) {
	i.offset = int(reader.ReadInt16())
}

func (i *BranchInstruction) Execute(frame *runtime.Frame) {
}

func (i *BranchInstruction) Branch(frame *runtime.Frame, offset int) {
	frame.SetNextPC(frame.GetThread().GetPC() + offset)
}

type Index8Instruction struct {
	index uint
}

func (i *Index8Instruction) FetchOperands(reader *runtime.BytecodeReader) {
	i.index = uint(reader.ReadUint8())
}

func (i *Index8Instruction) Execute(frame *runtime.Frame) {
}

type Index16Instruction struct {
	index uint
}

func (i *Index16Instruction) FetchOperands(reader *runtime.BytecodeReader) {
	i.index = uint(reader.ReadUint16())
}

func (i *Index16Instruction) Execute(frame *runtime.Frame) {
}
