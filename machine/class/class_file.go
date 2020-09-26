package class

type ClassFile struct {
	magic        uint32
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []*MemberInfo
	methods      []*MemberInfo
	attributes   []AttributeInfo
}

type MemberInfo struct {
	cp              ConstantPool
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributes      []AttributeInfo
}

func readMembers(reader *ClassFileReader, cp ConstantPool) []*MemberInfo {
	n := reader.readUint16()
	memberInfos := make([]*MemberInfo, n)
	for i := range memberInfos {
		memberInfos[i] = &MemberInfo{
			cp:              cp,
			accessFlags:     reader.readUint16(),
			nameIndex:       reader.readUint16(),
			descriptorIndex: reader.readUint16(),
			attributes:      readAttributes(reader, cp),
		}
	}
	return memberInfos
}

func (f *ClassFile) read(reader *ClassFileReader) error {
	f.readMagic(reader)
	f.readVersion(reader)
	f.readConstantPool(reader)
	f.readAccessFlags(reader)
	f.readClassAndInterfaces(reader)
	f.readFields(reader)
	f.readMethods(reader)
	f.readAttributes(reader)
	return nil
}

func (f *ClassFile) readMagic(reader *ClassFileReader) {
	magic := reader.readUint32()
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic!")
	}
	f.magic = magic
}

func (f *ClassFile) readVersion(reader *ClassFileReader) {
	minorVersion := reader.readUint16()
	majorVersion := reader.readUint16()
	if majorVersion < 42 || majorVersion > 52 {
		panic("java.lang.UnsupportedClassVersionError!")
	}
	f.minorVersion = minorVersion
	f.majorVersion = majorVersion
}

func (f *ClassFile) readConstantPool(reader *ClassFileReader) {
	n := int(reader.readUint16())
	cp := make([]ConstantInfo, n)
	for i := 1; i < n; i++ {
		cp[i] = newConstantInfo(reader, cp)
		switch cp[i].(type) {
		case *LongConstantInfo, *DoubleConstantInfo:
			i++
		}
	}
	f.constantPool = cp
}

func (f *ClassFile) readAccessFlags(reader *ClassFileReader) {
	accessFlags := reader.readUint16()
	f.accessFlags = accessFlags
}

func (f *ClassFile) readClassAndInterfaces(reader *ClassFileReader) {
	thisClass := reader.readUint16()
	superClass := reader.readUint16()
	interfaces := reader.readTable()
	f.thisClass = thisClass
	f.superClass = superClass
	f.interfaces = interfaces
}

func (f *ClassFile) readFields(reader *ClassFileReader) {
	f.fields = readMembers(reader, f.constantPool)
}

func (f *ClassFile) readMethods(reader *ClassFileReader) {
	f.methods = readMembers(reader, f.constantPool)
}

func (f *ClassFile) readAttributes(reader *ClassFileReader) {
	f.attributes = readAttributes(reader, f.constantPool)
}
