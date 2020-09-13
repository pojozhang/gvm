package class

type ConstantPool []ConstantInfo

func (p ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if c := p[index]; c != nil {
		return c
	}
	panic("Invalid constant pool index")
}

func (p ConstantPool) getUtf8(index uint16) string {
	return p.getConstantInfo(index).(*Utf8ConstantInfo).value
}
