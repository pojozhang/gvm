package control

import (
	"gvm/machine/runtime"
	"gvm/machine/runtime/instruction"
)

type GoTo struct {
	instruction.BranchInstruction
}

func (i *GoTo) Execute(frame *runtime.Frame) {
	frame.Jump(i.GetOffset())
}
