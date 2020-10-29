package control

import "gvm/machine/runtime"

type LookupSwitch struct {
	defaultOffset int32
	pairs         int32
	matchOffsets  []int32
}

func (i *LookupSwitch) FetchOperands(reader *runtime.BytecodeReader) {
	reader.SkipPadding()
	i.defaultOffset = reader.ReadInt32()
	i.pairs = reader.ReadInt32()
	i.matchOffsets = reader.ReadInt32s(i.pairs * 2)
}

func (i *LookupSwitch) Execute(frame *runtime.Frame) {
	key := frame.GetOperandStack().PopInt()
	for j := int32(0); j < i.pairs*2; j += 2 {
		if i.matchOffsets[j] == key {
			frame.Jump(int(i.matchOffsets[j+1]))
			return
		}
	}
	frame.Jump(int(i.defaultOffset))
}
