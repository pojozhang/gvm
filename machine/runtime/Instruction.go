package runtime

type Instruction interface {
	FetchOperands(reader *BytecodeReader)
	Execute(frame *Frame)
}

type BytecodeReader struct {
	pc   int
	code []byte
}

func (r *BytecodeReader) ReadUint8() uint8 {
	b := r.code[r.pc]
	r.pc++
	return b
}

func (r *BytecodeReader) ReadInt8() int8 {
	return int8(r.ReadUint8())
}

func (r *BytecodeReader) ReadUint16() uint16 {
	return (uint16(r.ReadUint8()) << 8) | uint16(r.ReadUint8())
}

func (r *BytecodeReader) ReadInt16() int16 {
	return int16(r.ReadUint8())
}

func (r *BytecodeReader) ReadInt32() int32 {
	return (int32(r.ReadUint8()) << 24) | (int32(r.ReadUint8()) << 16) | (int32(r.ReadUint8()) << 8) | int32(r.ReadUint8())
}

func (r *BytecodeReader) ReadInt32s(n int32) []int32 {
	int32s := make([]int32, n)
	for i := range int32s {
		int32s[i] = r.ReadInt32()
	}
	return int32s
}

func (r *BytecodeReader) SkipPadding() {
	for r.pc%4 != 0 {
		r.ReadUint8()
	}
}
