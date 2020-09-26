package class

type AttributeInfo interface {
	read(reader *ClassFileReader)
}

func readAttributes(reader *ClassFileReader, cp ConstantPool) []AttributeInfo {
	n := int(reader.readUint16())
	attributes := make([]AttributeInfo, n)
	for i := range attributes {
		attributes[i] = newAttributeInfo(reader, cp)
	}
	return attributes
}

func newAttributeInfo(reader *ClassFileReader, cp ConstantPool) AttributeInfo {
	name := cp.getUtf8(reader.readUint16())
	length := reader.readUint32()
	var attributeInfo AttributeInfo
	switch name {
	case "Code":
		attributeInfo = &CodeAttributeInfo{cp: cp}
	case "ConstantValue":
		attributeInfo = &ConstantValueAttributeInfo{}
	case "Deprecated":
		attributeInfo = &DeprecatedAttributeInfo{}
	case "Exceptions":
		attributeInfo = &ExceptionAttributionInfo{}
	case "LineNumberTable":
		attributeInfo = &LineNumberTableAttributeInfo{}
	case "LocalVariableTable":
		attributeInfo = &LocalVariableTableAttributeInfo{}
	case "SourceFile":
		attributeInfo = &SourceFileAttributeInfo{cp: cp}
	case "Synthetic":
		attributeInfo = &SyntheticAttributeInfo{}
	default:
		attributeInfo = &UnparsedAttributeInfo{name: name, length: length}
	}
	attributeInfo.read(reader)
	return attributeInfo
}

type MarkerAttributeInfo struct {
}

func (a *MarkerAttributeInfo) read(reader *ClassFileReader) {
}

type DeprecatedAttributeInfo struct {
	MarkerAttributeInfo
}

type SyntheticAttributeInfo struct {
	MarkerAttributeInfo
}

type SourceFileAttributeInfo struct {
	cp    ConstantPool
	index uint16
}

func (a *SourceFileAttributeInfo) read(reader *ClassFileReader) {
	a.index = reader.readUint16()
}

type ConstantValueAttributeInfo struct {
	index uint16
}

func (a *ConstantValueAttributeInfo) read(reader *ClassFileReader) {
	a.index = reader.readUint16()
}

type CodeAttributeInfo struct {
	cp             ConstantPool
	maxStack       uint16
	maxLocals      uint16
	code           []byte
	exceptionTable []*ExceptionTableEntry
	attributes     []AttributeInfo
}

func (a *CodeAttributeInfo) read(reader *ClassFileReader) {
	a.maxStack = reader.readUint16()
	a.maxLocals = reader.readUint16()
	a.code = reader.readBytes(reader.readUint32())
	a.exceptionTable = readExceptionTable(reader)
	a.attributes = readAttributes(reader, a.cp)
}

func readExceptionTable(reader *ClassFileReader) []*ExceptionTableEntry {
	length := reader.readUint16()
	exceptionTable := make([]*ExceptionTableEntry, length)
	for i := range exceptionTable {
		exceptionTable[i] = &ExceptionTableEntry{
			startPc:   reader.readUint16(),
			endPc:     reader.readUint16(),
			handlerPc: reader.readUint16(),
			catchType: reader.readUint16(),
		}
	}
	return exceptionTable
}

type ExceptionTableEntry struct {
	startPc   uint16
	endPc     uint16
	handlerPc uint16
	catchType uint16
}

type ExceptionAttributionInfo struct {
	exceptionIndexTable []uint16
}

func (a *ExceptionAttributionInfo) read(reader *ClassFileReader) {
	a.exceptionIndexTable = reader.readTable()
}

type LineNumberTableAttributeInfo struct {
	lineNumberTable []*LineNumberTableEntry
}

func (a *LineNumberTableAttributeInfo) read(reader *ClassFileReader) {
	length := reader.readUint16()
	lineNumberTable := make([]*LineNumberTableEntry, length)
	for i := range lineNumberTable {
		lineNumberTable[i] = &LineNumberTableEntry{
			startPc:    reader.readUint16(),
			lineNumber: reader.readUint16(),
		}
	}
	a.lineNumberTable = lineNumberTable
}

type LineNumberTableEntry struct {
	startPc    uint16
	lineNumber uint16
}

type LocalVariableTableAttributeInfo struct {
	localVariableTable []*LocalVariableTableEntry
}

func (a *LocalVariableTableAttributeInfo) read(reader *ClassFileReader) {
	length := reader.readUint16()
	localVariableTable := make([]*LocalVariableTableEntry, length)
	for i := range localVariableTable {
		localVariableTable[i] = &LocalVariableTableEntry{
			startPc:         reader.readUint16(),
			length:          reader.readUint16(),
			nameIndex:       reader.readUint16(),
			descriptorIndex: reader.readUint16(),
			index:           reader.readUint16(),
		}
	}
	a.localVariableTable = localVariableTable
}

type LocalVariableTableEntry struct {
	startPc         uint16
	length          uint16
	nameIndex       uint16
	descriptorIndex uint16
	index           uint16
}

type UnparsedAttributeInfo struct {
	name   string
	length uint32
	value  []byte
}

func (a *UnparsedAttributeInfo) read(reader *ClassFileReader) {
	a.value = reader.readBytes(a.length)
}
