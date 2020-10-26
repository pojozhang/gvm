package instruction

import "gvm/machine/runtime"

type AConstNull struct {
	NoOperandsInstruction
}

func (i AConstNull) Execute(frame *runtime.Frame) {
	frame.GetOperandStack().PushRef(nil)
}

type DConst0 struct {
	NoOperandsInstruction
}

func (i DConst0) Execute(frame *runtime.Frame) {
	frame.GetOperandStack().PushDouble(0)
}

type DConst1 struct {
	NoOperandsInstruction
}

func (i DConst1) Execute(frame *runtime.Frame) {
	frame.GetOperandStack().PushDouble(1)
}

type FConst0 struct {
	NoOperandsInstruction
}

func (i FConst0) Execute(frame *runtime.Frame) {
	frame.GetOperandStack().PushFloat(0)
}

type FConst1 struct {
	NoOperandsInstruction
}

func (i FConst1) Execute(frame *runtime.Frame) {
	frame.GetOperandStack().PushFloat(1)
}

type FConst2 struct {
	NoOperandsInstruction
}

func (i FConst2) Execute(frame *runtime.Frame) {
	frame.GetOperandStack().PushFloat(2)
}

type IConstM1 struct {
	NoOperandsInstruction
}

func (i IConstM1) Execute(frame *runtime.Frame) {
	frame.GetOperandStack().PushInt(-1)
}

type IConst0 struct {
	NoOperandsInstruction
}

func (i IConst0) Execute(frame *runtime.Frame) {
	frame.GetOperandStack().PushInt(0)
}

type IConst1 struct {
	NoOperandsInstruction
}

func (i IConst1) Execute(frame *runtime.Frame) {
	frame.GetOperandStack().PushInt(1)
}

type IConst2 struct {
	NoOperandsInstruction
}

func (i IConst2) Execute(frame *runtime.Frame) {
	frame.GetOperandStack().PushInt(2)
}

type IConst3 struct {
	NoOperandsInstruction
}

func (i IConst3) Execute(frame *runtime.Frame) {
	frame.GetOperandStack().PushInt(3)
}

type IConst4 struct {
	NoOperandsInstruction
}

func (i IConst4) Execute(frame *runtime.Frame) {
	frame.GetOperandStack().PushInt(4)
}

type IConst5 struct {
	NoOperandsInstruction
}

func (i IConst5) Execute(frame *runtime.Frame) {
	frame.GetOperandStack().PushInt(5)
}

type LConst0 struct {
	NoOperandsInstruction
}

func (i LConst0) Execute(frame *runtime.Frame) {
	frame.GetOperandStack().PushLong(0)
}

type LConst1 struct {
	NoOperandsInstruction
}

func (i LConst1) Execute(frame *runtime.Frame) {
	frame.GetOperandStack().PushLong(1)
}

type BiPush struct {
	value int8
}

func (i *BiPush) FetchOperands(reader *runtime.BytecodeReader) {
	i.value = reader.ReadInt8()
}

func (i *BiPush) Execute(frame *runtime.Frame) {
	frame.GetOperandStack().PushInt(int32(i.value))
}

type SiPush struct {
	value int16
}

func (i *SiPush) FetchOperands(reader *runtime.BytecodeReader) {
	i.value = reader.ReadInt16()
}

func (i *SiPush) Execute(frame *runtime.Frame) {
	frame.GetOperandStack().PushInt(int32(i.value))
}
