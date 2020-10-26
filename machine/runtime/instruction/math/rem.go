package math

import (
	"gvm/machine/runtime"
	"gvm/machine/runtime/instruction"
	"math"
)

type IRem struct {
	instruction.NoOperandsInstruction
}

func (i *IRem) Execute(frame *runtime.Frame) {
	stack := frame.GetOperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopInt()
	if v1 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	stack.PushInt(v2 % v1)
}

type LRem struct {
	instruction.NoOperandsInstruction
}

func (i *LRem) Execute(frame *runtime.Frame) {
	stack := frame.GetOperandStack()
	v1 := stack.PopLong()
	v2 := stack.PopLong()
	if v1 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	stack.PushLong(v2 % v1)
}

type FRem struct {
	instruction.NoOperandsInstruction
}

func (i *FRem) Execute(frame *runtime.Frame) {
	stack := frame.GetOperandStack()
	v1 := stack.PopFloat()
	v2 := stack.PopFloat()
	stack.PushDouble(math.Mod(float64(v2), float64(v1)))
}

type DRem struct {
	instruction.NoOperandsInstruction
}

func (i *DRem) Execute(frame *runtime.Frame) {
	stack := frame.GetOperandStack()
	v1 := stack.PopDouble()
	v2 := stack.PopDouble()
	stack.PushDouble(math.Mod(v2, v1))
}
