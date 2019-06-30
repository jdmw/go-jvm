package classloader

import (
	"os"
	"path/filepath"
	"strings"
)

type ClassPath struct{
	bootstrapClassPath Entry
	extClassPath Entry
	userClassPath Entry
}

func NewClassPath(jrePath string,cp string) *ClassPath  {
	classpath := &ClassPath{}

	jrePath = getJrePath(jrePath)

	classpath.bootstrapClassPath = newWildcardEntry(filepath.Join(jrePath,"lib","*"))
	classpath.extClassPath = newWildcardEntry(filepath.Join(jrePath,"lib","ext","*"))

	if cp == "" {
		cp = "/"
	}

	classpath.userClassPath = newEntry(cp)
	return classpath
}

func (self *ClassPath) ReadClass(classname string) ([]byte,Entry,error){
	classname = strings.Replace(classname,".","/",-1) + ".class"

	if data,entry,err := self.bootstrapClassPath.readClass(classname) ; err == nil {
		return data,entry,nil
	}
	if data,entry,err := self.extClassPath.readClass(classname) ; err == nil {
		return data,entry,nil
	}
	return self.userClassPath.readClass(classname)
}

func getJrePath(jrePath string) string{
	if jrePath != "" && exists(jrePath) {
		return jrePath
	}

	if exists("./jre") {
		return "./jre"
	}

	if home := os.Getenv("JAVA_HOME");home != "" {
		jrePath := filepath.Join(home,"jre")
		if exists(jrePath) {
			return jrePath
		}
	}

	panic("can't find jre")
}

func exists(path string) bool{
	if _,err := os.Stat(path) ;err != nil {
		if os.IsNotExist(err) {
			return false ;
		}
	}
	return true
}

func (self *ClassPath) String() string {
	return self.userClassPath.String()
}