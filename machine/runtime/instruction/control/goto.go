package control

import (
	"gvm/machine/runtime"
	"gvm/machine/runtime/instruction"
)

type GoTo struct {
	instruction.JumpInstruction
}

func (i *GoTo) Execute(frame *runtime.Frame) {
	frame.Jump(i.GetOffset())
}

type WGoTo struct {
	offset int
}

func (i *WGoTo) FetchOperands(reader *runtime.BytecodeReader) {
	i.offset = int(reader.ReadInt32())
}

func (i *WGoTo) Execute(frame *runtime.Frame) {
	frame.Jump(i.offset)
}
