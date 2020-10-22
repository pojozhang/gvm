package math

import (
	"gvm/machine/runtime"
	instruction "gvm/machine/runtime/instructions"
)

type IShl struct {
	instruction.NoOperandsInstruction
}

func (i *IShl) Execute(frame *runtime.Frame) {
	stack := frame.GetOperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopInt()
	stack.PushInt(v2 << (uint32(v1) & 0x1f))
}

type IShr struct {
	instruction.NoOperandsInstruction
}

func (i *IShr) Execute(frame *runtime.Frame) {
	stack := frame.GetOperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopInt()
	stack.PushInt(v2 >> (uint32(v1) & 0x1f))
}

type IUShr struct {
	instruction.NoOperandsInstruction
}

func (i *IUShr) Execute(frame *runtime.Frame) {
	stack := frame.GetOperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopInt()
	stack.PushInt(int32(uint32(v2) >> (uint32(v1) & 0x1f)))
}

type LShl struct {
	instruction.NoOperandsInstruction
}

func (i *LShl) Execute(frame *runtime.Frame) {
	stack := frame.GetOperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopLong()
	stack.PushLong(v2 << (uint32(v1) & 0x3f))
}

type LShr struct {
	instruction.NoOperandsInstruction
}

func (i *LShr) Execute(frame *runtime.Frame) {
	stack := frame.GetOperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopLong()
	stack.PushLong(v2 >> (uint32(v1) & 0x3f))
}

type LUShr struct {
	instruction.NoOperandsInstruction
}

func (i *LUShr) Execute(frame *runtime.Frame) {
	stack := frame.GetOperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopLong()
	stack.PushLong(int64(uint64(v2) >> (uint32(v1) & 0x3f)))
}
