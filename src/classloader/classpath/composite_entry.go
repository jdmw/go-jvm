package classloader

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

type CompositeEntry []Entry

func newCompositeEntry(path string) CompositeEntry  {
	paths := strings.Split(path,pathListSeparator)
	compositeEntry := make([]Entry,len(paths))
	for i,p := range paths {
		compositeEntry[i] = newEntry(p)
	}
	return CompositeEntry(compositeEntry)
}

func (self CompositeEntry) readClass(classname string) ([]byte,Entry,error) {
	for _,entry := range self {
		data ,from,err := entry.readClass(classname)
		if err == nil {
			return data,from,nil
		}
	}
	return nil,nil,errors.New("class not found: " + classname)
}

func (self CompositeEntry) String() string {
	strs := make([]string,len(self))
	for i,entry := range self {
		strs[i] = entry.String()
	}
	return strings.Join(strs,pathListSeparator)
}


func newWildcardEntry(path string) CompositeEntry {
	compositeEntry := []Entry{}
	baseDir := path[:len(path)-1]

	walkFn := func(f string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() && f != baseDir {
			return filepath.SkipDir ;
		}

		if strings.HasSuffix(f,".jar") || strings.HasSuffix(f,".JAR") {
			compositeEntry= append(compositeEntry,newJarEntry(f))
		}

		return nil
	}
	filepath.Walk(baseDir,walkFn)
	return compositeEntry

}