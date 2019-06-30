package classloader

import (
	"os"
	"strings"
)

const pathListSeparator = string(os.PathListSeparator)
type Entry interface {
	readClass(classname string) ([]byte,Entry,error)
	String() string
}

func newEntry(path string) Entry {

	if strings.Contains(path,pathListSeparator) {
		return newCompositeEntry(path);
	}
	if strings.HasSuffix(path,"*"){
		return newWildcardEntry(path)
	}
	if strings.HasSuffix(path,".jar") || strings.HasSuffix(path,".JAR") {
		return newJarEntry(path);
	}
	return newDirEntry(path);
}