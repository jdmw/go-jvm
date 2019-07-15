package classloader

import (
	"./classfile"
	"./classpath"
	"fmt"
)

type ClassLoder struct {
	classPath classpath.ClassPath
	loadedClasses map[string]string
}

func (self ClassLoder) LoadClass(classname string) *Class{

	data,_,_ := self.classPath.ReadClass(classname)
	if data == nil {
		//fmt.Errorf("can not load main class " + classname)
		return  nil
	} else {
		classfile := classfile.ParseClassFile(data)
		return self.buildClass(classfile)
	}
}


func NewClassloader(xjre string,cp string)  ClassLoder{
	return ClassLoder{*classpath.NewClassPath(xjre,cp),make(map[string]string)}
}

