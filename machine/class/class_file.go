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
	attributes   []*AttributeInfo
}

type MemberInfo struct {
	cp              ConstantPool
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributes      []AttributeInfo
}

func (f *ClassFile) read(r *ClassFileReader) error {
	f.readMagic(r)
	f.readVersion(r)
	f.readConstantPool(r)
	f.readAccessFlags(r)
	f.readClassAndInterfaces(r)
	return nil
}

func (f *ClassFile) readMagic(r *ClassFileReader) {
	magic := r.readUint32()
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic!")
	}
	f.magic = magic
}

func (f *ClassFile) readVersion(r *ClassFileReader) {
	minorVersion := r.readUint16()
	majorVersion := r.readUint16()
	if majorVersion < 42 || majorVersion > 52 {
		panic("java.lang.UnsupportedClassVersionError!")
	}
	f.minorVersion = minorVersion
	f.majorVersion = majorVersion
}

func (f *ClassFile) readConstantPool(r *ClassFileReader) {
	n := int(r.readUint16())
	cp := make([]ConstantInfo, n)
	for i := 1; i < n; i++ {
		cp[i] = newConstantInfo(r, cp)
		switch cp[i].(type) {
		case *LongConstantInfo, *DoubleConstantInfo:
			i++
		}
	}
	f.constantPool = cp
}

func (f *ClassFile) readAccessFlags(r *ClassFileReader) {
	accessFlags := r.readUint16()
	f.accessFlags = accessFlags
}

func (f *ClassFile) readClassAndInterfaces(r *ClassFileReader) {
	thisClass := r.readUint16()
	superClass := r.readUint16()
	interfaces := r.readTable()
	f.thisClass = thisClass
	f.superClass = superClass
	f.interfaces = interfaces
}

func parseMemberInfos(f *ClassFile, r *ClassFileReader) []*MemberInfo {
	n := r.readUint16()
	memberInfos := make([]*MemberInfo, n)
	for i := range memberInfos {
		memberInfos[i] = &MemberInfo{
			cp:              f.constantPool,
			accessFlags:     r.readUint16(),
			nameIndex:       r.readUint16(),
			descriptorIndex: r.readUint16(),
			attributes:      nil,
		}
	}
	return memberInfos
}
