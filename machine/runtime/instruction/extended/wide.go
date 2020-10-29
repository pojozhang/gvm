package extended

import (
	"gvm/machine/runtime"
	"gvm/machine/runtime/instruction/load"
	"gvm/machine/runtime/instruction/math"
	"gvm/machine/runtime/instruction/store"
)

// Wide指令的功能是改变其他指令。
type Wide struct {
	delegate runtime.Instruction
}

func (i *Wide) FetchOperands(reader *runtime.BytecodeReader) {
	code := reader.ReadUint8()
	switch code {
	case 0x15:
		i.delegate = &load.WILoad{}
	case 0x16:
		i.delegate = &load.WLLoad{}
	case 0x17:
		i.delegate = &load.WFLoad{}
	case 0x18:
		i.delegate = &load.WDLoad{}
	case 0x19:
		i.delegate = &load.WALoad{}
	case 0x36:
		i.delegate = &store.WIStore{}
	case 0x37:
		i.delegate = &store.WLStore{}
	case 0x38:
		i.delegate = &store.WFStore{}
	case 0x39:
		i.delegate = &store.WDStore{}
	case 0x3a:
		i.delegate = &store.WAStore{}
	case 0x84:
		i.delegate = &math.WIInc{}
	case 0xa9:
		panic("Unsupported opcode: 0xa9!")
	}
}

func (i *Wide) Execute(frame *runtime.Frame) {
	i.delegate.Execute(frame)
}
