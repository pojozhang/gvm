package class

import "math"

const (
	ConstantUtf8               = 1
	ConstantInteger            = 3
	ConstantFloat              = 4
	ConstantLong               = 5
	ConstantDouble             = 6
	ConstantClass              = 7
	ConstantString             = 8
	ConstantFieldRef           = 9
	ConstantMethodRef          = 10
	ConstantInterfaceMethodRef = 11
	ConstantNameAndType        = 12
	ConstantMethodHandle       = 15
	ConstantMethodType         = 16
	ConstantInvokeDynamic      = 18
)

type ConstantInfo interface {
	read(r *ClassFileReader)
}

func newConstantInfo(r *ClassFileReader, cp ConstantPool) ConstantInfo {
	tag := r.readUint8()
	var c ConstantInfo
	switch tag {
	case ConstantUtf8:
		c = &Utf8ConstantInfo{}
	case ConstantInteger:
		c = &IntegerConstantInfo{}
	case ConstantLong:
		c = &LongConstantInfo{}
	case ConstantFloat:
		c = &FloatConstantInfo{}
	case ConstantDouble:
		c = &DoubleConstantInfo{}
	case ConstantClass:
		c = &ClassConstantInfo{cp: cp}
	case ConstantString:
		c = &StringConstantInfo{cp: cp}
	case ConstantFieldRef:
		c = &FieldRefConstantInfo{}
	case ConstantMethodRef:
		c = &MethodRefConstantInfo{}
	case ConstantInterfaceMethodRef:
		c = &InterfaceMethodRefConstantInfo{}
	case ConstantNameAndType:
		c = &NameAndTypeConstantInfo{cp: cp}
	case ConstantMethodHandle:
	case ConstantMethodType:
	case ConstantInvokeDynamic:
	}
	c.read(r)
	return c
}

type IntegerConstantInfo struct {
	value int32
}

func (c *IntegerConstantInfo) read(r *ClassFileReader) {
	c.value = int32(r.readUint32())
}

type LongConstantInfo struct {
	value int64
}

func (c *LongConstantInfo) read(r *ClassFileReader) {
	c.value = int64(r.readUint64())
}

type FloatConstantInfo struct {
	value float32
}

func (c *FloatConstantInfo) read(r *ClassFileReader) {
	c.value = math.Float32frombits(r.readUint32())
}

type DoubleConstantInfo struct {
	value float64
}

func (c *DoubleConstantInfo) read(r *ClassFileReader) {
	c.value = math.Float64frombits(r.readUint64())
}

type Utf8ConstantInfo struct {
	value string
}

func (c *Utf8ConstantInfo) read(r *ClassFileReader) {
	length := r.readUint16()
	bytes := r.readBytes(uint32(length))
	c.value = string(bytes)
}

type StringConstantInfo struct {
	cp        ConstantPool
	utf8Index uint16
}

func (c *StringConstantInfo) read(r *ClassFileReader) {
	c.utf8Index = r.readUint16()
}

type ClassConstantInfo struct {
	cp        ConstantPool
	nameIndex uint16
}

func (c *ClassConstantInfo) read(r *ClassFileReader) {
	c.nameIndex = r.readUint16()
}

type NameAndTypeConstantInfo struct {
	cp              ConstantPool
	nameIndex       uint16
	descriptorIndex uint16
}

func (c *NameAndTypeConstantInfo) read(r *ClassFileReader) {
	c.nameIndex = r.readUint16()
	c.descriptorIndex = r.readUint16()
}

type MemberRefConstantInfo struct {
	cp               ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}

func (c *MemberRefConstantInfo) read(r *ClassFileReader) {
	c.classIndex = r.readUint16()
	c.nameAndTypeIndex = r.readUint16()
}

type FieldRefConstantInfo struct {
	MemberRefConstantInfo
}

type MethodRefConstantInfo struct {
	MemberRefConstantInfo
}

type InterfaceMethodRefConstantInfo struct {
	MemberRefConstantInfo
}
