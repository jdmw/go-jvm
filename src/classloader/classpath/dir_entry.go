package classpath
import (
	"io/ioutil"
	"path/filepath"
)

type DirEntry struct {
	dir string
}

func newDirEntry (path string) *DirEntry{
	absDir,error := filepath.Abs(path)
	if error != nil {
		panic(error)
	}
	return &DirEntry{absDir}
}

func (self *DirEntry) readClass(classname string) ([]byte,Entry,error) {
	data,err := ioutil.ReadFile(filepath.Join(self.dir,classname))
	return data,self,err
}

func (self *DirEntry) String() string {
	return self.dir;
}
