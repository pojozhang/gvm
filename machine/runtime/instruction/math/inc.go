package math

import (
	"gvm/machine/runtime"
)

type IInc struct {
	index     uint
	increment int32
}

func (i *IInc) FetchOperands(reader *runtime.BytecodeReader) {
	i.index, i.increment = uint(reader.ReadUint8()), int32(reader.ReadInt8())
}

func (i *IInc) Execute(frame *runtime.Frame) {
	localVars := frame.GetLocalVars()
	localVars.SetInt(i.index, localVars.GetInt(i.index)+i.increment)
}
