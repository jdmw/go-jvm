package classloader

import (
	"./classfile"
	"strings"

	"../util"
)

type Class struct {

	AccFlag uint16
	Inited bool
	Classloader ClassLoder
	Classname string
	ParentClass *Class
	Interfaces []*Class
	Constantpool ConstantPool
	Fields []Field
	Methods []Method
	Attributes classfile.Attributes
}

func (self ClassLoder) buildClass(cf *classfile.ClassFile ) *Class {
	class := Class{}
	class.Classloader = self
	accFlag,thisclass,superclass,interfacenames,cpool,fields,methods,attributes := cf.ClassInfo()
	self.loadedClasses[thisclass] = &class
	class.AccFlag = accFlag
	class.Classname = thisclass
	class.ParentClass = self.LoadClass(superclass)
	class.Interfaces = make([]*Class,len(interfacenames))
	for i,name := range interfacenames {
		class.Interfaces[i] = self.LoadClass(name)
	}
	class.Constantpool = ConstantPool{cpool,make([]*util.Slot,len(cpool))}
	class.Fields = make([]Field,len(interfacenames))
	for i,fieldinfo := range fields {
		class.Fields[i] = Field{&class,fieldinfo.AccessFlag(),fieldinfo.Name(),
			fieldinfo.Descriptor(),fieldinfo.Attributes()}
	}
	class.Methods = make([]Method,len(methods))
	for i,info := range methods {
		class.Methods[i] = Method{&class,info.AccessFlag(),info.Name(),info.Descriptor(),
			info.Attributes(),false,0,0,nil }
	}
	class.Attributes = attributes
	return &class
}

func (self Class) LoadConstant (index util.U2,slotTable util.SlotTable) {
	//entry := self.constantpool.slots[index]
	//if( entry == nil){
		raw := self.Constantpool.cpool[index]
		switch raw.(type) {
			//case *classfile.ConstClassInfo :
			//	clazz := self.classloader.LoadClass(raw.(*classfile.ConstClassInfo).Name())
			//	slotTable.SetObjectRef(index,*clazz)
			//case *classfile.ConstClassInfo :
			//case *ConstMemberRefInfo{cp,0,0};break;
			//case *ConstMemberRefInfo{cp,0,0};break;
			//case *ConstMemberRefInfo{cp,0,0};break;
			//case *ConstStringInfo{cp,0};break;
			case *classfile.ConstIntegerInfo:
				slotTable.SetU4(index,util.U4(raw.(*classfile.ConstIntegerInfo).Value()))
				break
			case *classfile.ConstFloatInfo:
				slotTable.SetFloat(index,raw.(*classfile.ConstFloatInfo).Value())
				break
			case *classfile.ConstLongInfo:
				slotTable.SetU8(index,raw.(*classfile.ConstLongInfo).Value())
				break;
			case *classfile.ConstDoubleInfo:
				slotTable.SetDouble(index,raw.(*classfile.ConstDoubleInfo).Value())
				break;


			//case *ConstNameAndTypeInfo{};break;
			//case *ConstUtf8Info{};break;
			//case *ConstMethodHandleInfo{};break;
			//case *ConstMethodTypeInfo{};break;
			//case *ConstDynamicInfo{};break;
			//case *ConstInvokeDynamicInfo{};break;
			//case *ConstModuleInfo{};checkAccFlag = ACC_MODULE;break;
			//case *ConstPackageInfo{};checkAccFlag = ACC_MODULE;break;

		}
	//}
	//return entry
}


type ConstantPool struct {
	cpool classfile.ConstantPool
	slots []*util.Slot
}

/**
 * find method by name and parameter types
 * eg: find main method : FindMethod("main","([Ljava/lang/String;)V")
 */
func (self *Class) FindMethod(methodName string,descriptor string) *Method{
	if ! strings.HasSuffix(descriptor,")V" ) {
		descriptor = "([" + descriptor + ")V"
	}

	for _,method := range self.Methods {
		if(strings.Compare(method.Name , methodName) == 0 && strings.Compare(descriptor ,method.Descripter )==0){
			return &method ;
		}
	}
	return nil
}