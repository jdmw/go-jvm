package classloader


import (
	"./classfile"
	"./classpath"
	"fmt"
)

type Class struct {

	//AccFlag u2
	classname string
	parentClass Class
	interfaces []Class
	unyyh
	fields []Field
	methods []Method
	attributes []Attribute

}

func (self ClassLoder) buildClass(cf *classfile.ClassFile ) *Class {
	class := Class{}
	thisclass,superclass,interfacenames,cp,fields,methods,attributes := cf.ClassInfo()
	class.classname = thisclass
	class.parentClass = *self.LoadClass(superclass)
	class.interfaces = make([]Class,len(interfacenames))
	for i,name := range interfacenames {
		class.interfaces[i] = *self.LoadClass(name)
	}
	class.fields = make([]Field,len(interfacenames))
	for i,fieldinfo := range fields {
		class.fields[i] = Field{class,fieldinfo.Name(),fieldinfo.Descriptor(),
			*buildAttributes(class,fieldinfo.Attributes())}
	}
	class.methods = make([]Method,len(methods))
	for i,info := range methods {
		class.methods[i] = Method{class,info.Name(),info.Descriptor(),
			*buildAttributes(class,info.Attributes())}
	}
	class.attributes = *buildAttributes(class,attributes)
	return &class
}


type Method struct {
	class Class
	Name string
	Desripter string
	attributes []Attribute
}
type Field struct {
	class Class
	Name string
	Desripter string
	attributes []Attribute
}

type Attribute struct {
	class Class
	info classfile.AttributeInfo
}


func buildAttributes(class Class,infos []classfile.AttributeInfo) *([]Attribute){
	attributes := make( []Attribute,len(infos))
	for i,info := range infos {
		attributes[i] = Attribute{class,info}
	}
	return &attributes
}