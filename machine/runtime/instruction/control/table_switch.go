package control

import (
	"gvm/machine/runtime"
)

type TableSwitch struct {
	defaultOffset int32   // default分支字节码跳转偏移量。
	low           int32   // case取值范围下限。
	high          int32   // case取值范围上限。
	jumpOffsets   []int32 // 除default以外的分支的字节码跳转偏移量。
}

func (i *TableSwitch) FetchOperands(reader *runtime.BytecodeReader) {
	reader.SkipPadding()
	i.defaultOffset = reader.ReadInt32()
	i.low = reader.ReadInt32()
	i.high = reader.ReadInt32()
	i.jumpOffsets = reader.ReadInt32s(i.high - i.low + 1)
}

func (i *TableSwitch) Execute(frame *runtime.Frame) {
	if index := frame.GetOperandStack().PopInt(); index >= i.low && index <= i.high {
		frame.Jump(int(i.jumpOffsets[index-i.low]))
		return
	}
	frame.Jump(int(i.defaultOffset))
}
