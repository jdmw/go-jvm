package classloader

import (
	"./classfile"
	"../util"
)

type Method struct {
	class *Class
	AccFlag uint16
	Name string
	Descripter string
	Attributes *classfile.Attributes
	// load
	load bool
	MaxStack util.U2
	MaxLocals util.U2
	Code []byte
}

type Field struct {
	class *Class
	AccFlag uint16
	Name string
	Descripter string
	Attributes *classfile.Attributes
}

/*type Attribute struct {
	class *Class
	info *classfile.AttributeInfo
}

func buildAttributes(class *Class,infos []*classfile.AttributeInfo) *([]Attribute){
	attributes := make( []Attribute,len(infos))
	for i,info := range infos {
		attributes[i] = Attribute{class,info}
	}
	return &attributes
}
*/

func (self *Method) GetCodeAttribute() *classfile.CodeAttr{
	return self.Attributes.GetCodeAttr()
}

func (self *Method) Load()  {
	if !self.load {
		codeAttr := self.GetCodeAttribute()
		if(codeAttr != nil){
			self.MaxLocals,self.MaxStack,self.Code = codeAttr.MaxLocals(),codeAttr.MaxStack(),codeAttr.Code()
			self.load = true
		}
	}
}

func (self *Method) IsNative() bool {
	return classfile.ACC_NATIVE & self.AccFlag == 1
}

func (self *Method) GetClass() *Class {
	return self.class
}