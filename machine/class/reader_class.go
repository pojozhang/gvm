package class

import (
	"io/ioutil"
	"path/filepath"
)

type FileReader struct {
	dir string
}

func NewFileReader(dir string) *FileReader {
	return &FileReader{
		dir,
	}
}

func (r *FileReader) readClass(className string) ([]byte, error) {
	return ioutil.ReadFile(filepath.Join(r.dir, className))
}
