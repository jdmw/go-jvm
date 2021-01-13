package classloader

import (
	"./classfile"
	"./classpath"
	"strings"

	//"fmt"
)

type ClassLoder struct {
	classPath classpath.ClassPath
	loadedClasses map[string]*Class
}

//
func (self ClassLoder) LoadClass(classname string) *Class{

	classname = strings.Replace(classname,".","/",-1)
	loadedClass := self.loadedClasses[classname]
	if(loadedClass != nil){
		return loadedClass
	}

	data,_,_ := self.classPath.ReadClass(classname)
	if data == nil {
		//fmt.Errorf("can not load main class " + classname)
		return  nil
	} else {
		classfile := classfile.ParseClassFile(data)
		loadedClass = self.buildClass(classfile)
		self.loadedClasses[classname] = loadedClass
		return loadedClass
	}
}

func NewClassloader(xjre string,cp string)  ClassLoder{
	return ClassLoder{*classpath.NewClassPath(xjre,cp),make(map[string]*Class)}
}

