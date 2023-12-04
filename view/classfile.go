package view

import "fmt"

type ClassFile struct {
	Magic        uint32           `json:"magic"`
	MinorVersion uint16           `json:"minorVersion"`
	MajorVersion uint16           `json:"majorVersion"`
	ConstantPool ConstantPoolInfo `json:"constantPool"`
	accessFlags  uint16           `json:"accessFlags"`
	thisClass    uint16           `json:"thisClass"`
	superClass   uint16           `json:"superClass"`
	Interfaces   []uint16         `json:"interfaces"`
	Fields       []*MemberInfo    `json:"fields"`
	Methods      []*MemberInfo    `json:"methods"`
	Attributes   []AttributeInfo  `json:"attributes"`
}

func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok := r.(error)
			if !ok {
				fmt.Println(err)
				err = fmt.Errorf("%v\n", r)
			}
		}
	}()
	cr := &ClassReader{classData, 0}
	cf = &ClassFile{}
	cf.read(cr)
	return
}

func (self *ClassFile) read(cr *ClassReader) {
	self.readMagic(cr)
	self.readVersion(cr)
	self.readConstantPool(cr)
	self.readAccessFlags(cr)
	self.readClassName(cr)
	self.readSuperClassName(cr)
	self.readInterfaces(cr)
	self.readFields(cr)
	self.readMethods(cr)
	self.readAttributes(cr)
}

func (self *ClassFile) readMagic(cr *ClassReader) {
	self.Magic = cr.readUInt32()
	if self.Magic != 0xCAFEBABE {
		panic("不是有效的java文件")
	}
}

func (self *ClassFile) readVersion(cr *ClassReader) {
	self.readMinorVersion(cr)
	self.readMajorVersion(cr)
	fmt.Println("class file minorVersion", self.MinorVersion, "majorVersion", self.MajorVersion)

}

func (self *ClassFile) readMajorVersion(cr *ClassReader) {
	self.MajorVersion = cr.readUInt16()
}

func (self *ClassFile) readMinorVersion(cr *ClassReader) {
	self.MinorVersion = cr.readUInt16()
}

func (self *ClassFile) readConstantPool(cr *ClassReader) {

}

func (self *ClassFile) readAccessFlags(cr *ClassReader) {

}

func (self *ClassFile) readFields(cr *ClassReader) {

}

func (self *ClassFile) readMethods(cr *ClassReader) {

}

func (self *ClassFile) readClassName(cr *ClassReader) {

}

func (self *ClassFile) readSuperClassName(cr *ClassReader) {

}
func (self *ClassFile) readInterfaces(cr *ClassReader) {

}
func (self *ClassFile) readAttributes(cr *ClassReader) {

}

type ConstantPoolInfo struct {
}

type MemberInfo struct {
}
type AttributeInfo struct {
}
