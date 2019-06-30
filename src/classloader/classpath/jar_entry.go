package classloader

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)
type JarEntry struct {
	path string
}

func newJarEntry(path string) *JarEntry{
	absPath,err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &JarEntry{absPath}
}

func (self *JarEntry) readClass(classname string) ([]byte,Entry,error)  {
	z,err := zip.OpenReader(self.path)
	if err != nil {
		return nil,nil,err
	}
	defer z.Close()
	for _, file := range z.File {
		if file.Name == classname {
			rc,err := file.Open()
			if err != nil {
				return nil,nil,err
			}
			defer rc.Close()
			data,err := ioutil.ReadAll(rc)
			if err != nil {
				return nil,nil,err
			}
			return data,self,nil
		}
	}
	return nil,nil,errors.New("class not found: " + classname)
}

func (self *JarEntry) String() string {
	return self.path;
}
