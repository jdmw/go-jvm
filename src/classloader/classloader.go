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

func (self ClassLoder) LoadClass(classname string) *classfile.ClassFile{

	data,_,_ := self.classPath.ReadClass(classname)
	if data == nil {
		fmt.Errorf("can not load main class " + classname)
		return  nil
	} else {
		return classfile.ParseClassFile(data)
	}
}

func NewClassloader(xjre string,cp string)  ClassLoder{
	return ClassLoder{*classpath.NewClassPath(xjre,cp),make(map[string]string)}
}

