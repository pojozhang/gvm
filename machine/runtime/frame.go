package runtime

import "math"

type Frame struct {
	next         *Frame
	localVars    LocalVars
	operandStack *OperandStack
	thread       *Thread
}

func (f *Frame) GetOperandStack() *OperandStack {
	return f.operandStack
}

func (f *Frame) GetLocalVars() LocalVars {
	return f.localVars
}

func (f *Frame) GetThread() *Thread {
	return f.thread
}

func (f *Frame) SetNextPC(offset int) {

}

func NewFrame(thread *Thread, maxLocals, maxStack uint) *Frame {
	return &Frame{
		thread:       thread,
		localVars:    newLocalVars(maxStack),
		operandStack: newOperandStack(maxStack),
	}
}

type Slot struct {
	num int32
	ref *Object
}

type LocalVars []Slot

func (v LocalVars) SetInt(index uint, value int32) {
	v[index].num = value
}

func (v LocalVars) GetInt(index uint) int32 {
	return v[index].num
}

func (v LocalVars) SetLong(index uint, value int64) {
	v[index].num = int32(value)
	v[index+1].num = int32(value >> 32)
}

func (v LocalVars) GetLong(index uint) int64 {
	return int64(v[index].num) | int64(v[index+1].num)<<32
}

func (v LocalVars) SetFloat(index uint, value float32) {
	v[index].num = int32(math.Float32bits(value))
}

func (v LocalVars) GetFloat(index uint) float32 {
	return math.Float32frombits(uint32(v[index].num))
}

func (v LocalVars) SetDouble(index uint, value float64) {
	v.SetLong(index, int64(math.Float64bits(value)))
}

func (v LocalVars) GetDouble(index uint) float64 {
	return math.Float64frombits(uint64(v.GetLong(index)))
}

func (v LocalVars) SetRef(index uint, object *Object) {
	v[index].ref = object
}

func (v LocalVars) GetRef(index uint) *Object {
	return v[index].ref
}

func newLocalVars(maxLocals uint) LocalVars {
	return make([]Slot, maxLocals)
}

type OperandStack struct {
	size  uint
	slots []Slot
}

func newOperandStack(maxStack uint) *OperandStack {
	return &OperandStack{
		slots: make([]Slot, maxStack),
	}
}

func (s *OperandStack) PushInt(value int32) {
	s.slots[s.size].num = value
	s.size++
}

func (s *OperandStack) PopInt() int32 {
	s.size--
	return s.slots[s.size].num
}

func (s *OperandStack) PushLong(value int64) {
	s.slots[s.size].num = int32(value)
	s.slots[s.size+1].num = int32(value >> 32)
	s.size += 2
}

func (s *OperandStack) PopLong() int64 {
	s.size -= 2
	return int64(uint32(s.slots[s.size].num)) | (int64(uint32(s.slots[s.size+1].num)) << 32)
}

func (s *OperandStack) PushFloat(value float32) {
	s.slots[s.size].num = int32(math.Float32bits(value))
	s.size++
}

func (s *OperandStack) PopFloat() float32 {
	s.size--
	return math.Float32frombits(uint32(s.slots[s.size].num))
}

func (s *OperandStack) PushDouble(value float64) {
	s.PushLong(int64(math.Float64bits(value)))
}

func (s *OperandStack) PopDouble() float64 {
	return math.Float64frombits(uint64(s.PopLong()))
}

func (s *OperandStack) PushRef(object *Object) {
	s.slots[s.size].ref = object
	s.size++
}

func (s *OperandStack) PopRef() *Object {
	s.size--
	return s.slots[s.size].ref
}

func (s *OperandStack) PushSlot(slot Slot) {
	s.slots[s.size] = slot
	s.size++
}

func (s *OperandStack) PopSlot() Slot {
	s.size--
	return s.slots[s.size]
}
