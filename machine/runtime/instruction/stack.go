package instruction

import "gvm/machine/runtime"

type Pop struct {
	NoOperandsInstruction
}

func (i *Pop) Execute(frame *runtime.Frame) {
	frame.GetOperandStack().PopSlot()
}

type Pop2 struct {
	NoOperandsInstruction
}

func (i *Pop2) Execute(frame *runtime.Frame) {
	stack := frame.GetOperandStack()
	stack.PopSlot()
	stack.PopSlot()
}

type Swap struct {
	NoOperandsInstruction
}

func (i *Swap) Execute(frame *runtime.Frame) {
	stack := frame.GetOperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
}
