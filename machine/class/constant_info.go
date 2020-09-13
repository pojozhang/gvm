package class

import "math"

const (
	CONSTANT_UTF8                 = 1
	CONSTANT_INTEGER              = 3
	CONSTANT_FLOAT                = 4
	CONSTANT_LONG                 = 5
	CONSTANT_DOUBLE               = 6
	CONSTANT_CLASS                = 7
	CONSTANT_STRING               = 8
	CONSTANT_FIELD_REF            = 9
	CONSTANT_METHOD_REF           = 10
	CONSTANT_INTERFACE_METHOD_REF = 11
	CONSTANT_NAME_AND_TYPE        = 12
	CONSTANT_METHOD_HANDLE        = 15
	CONSTANT_METHOD_TYPE          = 16
	CONSTANT_INVOKE_DYNAMIC       = 18
)

type ConstantInfo interface {
	read(r *ClassFileReader)
}

func newConstantInfo(r *ClassFileReader, cp ConstantPool) ConstantInfo {
	tag := r.readUint8()
	var c ConstantInfo
	switch tag {
	case CONSTANT_UTF8:
		c = &Utf8ConstantInfo{}
	case CONSTANT_INTEGER:
		c = &IntegerConstantInfo{}
	case CONSTANT_LONG:
		c = &LongConstantInfo{}
	case CONSTANT_FLOAT:
		c = &FloatConstantInfo{}
	case CONSTANT_DOUBLE:
		c = &DoubleConstantInfo{}
	case CONSTANT_CLASS:
		c = &ClassConstantInfo{cp: cp}
	case CONSTANT_STRING:
		c = &StringConstantInfo{cp: cp}
	case CONSTANT_FIELD_REF:
		c = &FieldRefConstantInfo{}
	case CONSTANT_METHOD_REF:
		c = &MethodRefConstantInfo{}
	case CONSTANT_INTERFACE_METHOD_REF:
		c = &InterfaceMethodRefConstantInfo{}
	case CONSTANT_NAME_AND_TYPE:
		c = &NameAndTypeConstantInfo{cp: cp}
	case CONSTANT_METHOD_HANDLE:
	case CONSTANT_METHOD_TYPE:
	case CONSTANT_INVOKE_DYNAMIC:
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
