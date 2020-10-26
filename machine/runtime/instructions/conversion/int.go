package conversion

import (
	"gvm/machine/runtime"
	instruction "gvm/machine/runtime/instructions"
)

type I2B struct {
	instruction.NoOperandsInstruction
}

func (i *I2B) Execute(frame *runtime.Frame) {
	stack := frame.GetOperandStack()
	stack.PushInt(int32(int8(stack.PopInt())))
}

type I2C struct {
	instruction.NoOperandsInstruction
}

func (i *I2C) Execute(frame *runtime.Frame) {
	stack := frame.GetOperandStack()
	stack.PushInt(int32(uint16(stack.PopInt())))
}

type I2S struct {
	instruction.NoOperandsInstruction
}

func (i *I2S) Execute(frame *runtime.Frame) {
	stack := frame.GetOperandStack()
	stack.PushInt(int32(int16(stack.PopInt())))
}

type I2L struct {
	instruction.NoOperandsInstruction
}

func (i *I2L) Execute(frame *runtime.Frame) {
	stack := frame.GetOperandStack()
	stack.PushLong(int64(stack.PopInt()))
}

type I2F struct {
	instruction.NoOperandsInstruction
}

func (i *I2F) Execute(frame *runtime.Frame) {
	stack := frame.GetOperandStack()
	stack.PushFloat(float32(stack.PopInt()))
}

type I2D struct {
	instruction.NoOperandsInstruction
}

func (i *I2D) Execute(frame *runtime.Frame) {
	stack := frame.GetOperandStack()
	stack.PushDouble(float64(stack.PopInt()))
}
