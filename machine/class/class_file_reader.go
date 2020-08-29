package class

import "encoding/binary"

type ClassFileReader struct {
	data []byte
}

func (r *ClassFileReader) readUint8() uint8 {
	return r.readBytes(1)[0]
}

func (r *ClassFileReader) readUint16() uint16 {
	return binary.BigEndian.Uint16(r.readBytes(2))
}

func (r *ClassFileReader) readUint32() uint32 {
	return binary.BigEndian.Uint32(r.readBytes(4))
}

func (r *ClassFileReader) readUint64() uint64 {
	return binary.BigEndian.Uint64(r.readBytes(8))
}

func (r *ClassFileReader) readBytes(n uint32) []byte {
	bytes := r.data[:n]
	r.data = r.data[n:]
	return bytes
}
