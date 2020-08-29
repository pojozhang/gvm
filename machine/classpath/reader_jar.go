package classpath

import (
	"archive/zip"
	"io/ioutil"
	"path/filepath"
)

type JarFileReader struct {
	path string
}

func NewJarFileReader(path string) (*JarFileReader, error) {
	p, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}
	return &JarFileReader{
		p,
	}, nil
}

func (r *JarFileReader) readClass(className string) ([]byte, error) {
	rc, err := zip.OpenReader(r.path)
	if err != nil {
		return nil, err
	}
	defer rc.Close()

	for _, file := range rc.File {
		rc, err := file.Open()
		if err != nil {
			return nil, err
		}
		defer rc.Close()
		println(file.Name)
		data, err := ioutil.ReadAll(rc)
		if err != nil {
			return nil, err
		}
		println(data)
	}
	return nil, nil
}
