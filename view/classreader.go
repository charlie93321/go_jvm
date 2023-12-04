package view

import "encoding/binary"

type ClassReader struct {
	data []byte
	pos  uint64
}

func (self *ClassReader) readUInt8() uint8 {
	b1 := self.data[self.pos]
	self.pos += 1
	return b1
}

func (self *ClassReader) readUInt16() uint16 {
	b2 := binary.BigEndian.Uint16(self.data[self.pos : self.pos+2])
	self.pos += 2
	return b2
}

func (self *ClassReader) readUInt32() uint32 {
	b4 := binary.BigEndian.Uint32(self.data[self.pos : self.pos+4])
	self.pos += 4
	return b4
}

func (self *ClassReader) readUInt64() uint64 {
	b8 := binary.BigEndian.Uint64(self.data[self.pos : self.pos+8])
	self.pos += 8
	return b8
}

func (self *ClassReader) readUInt16Arr() []uint16 {
	n := self.readUInt16()
	arr := make([]uint16, n)
	for i := range arr {
		arr[i] = self.readUInt16()
	}
	return arr
}

func (self *ClassReader) readBytes(n uint32) []byte {
	var bytes []uint8
	copy(self.data[:n], bytes)
	return bytes
}
