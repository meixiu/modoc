package helper

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
)

func ListFiles(dirname string) []string {
	f, _ := os.Open(dirname)
	names, _ := f.Readdirnames(-1)
	defer f.Close()
	sort.Strings(names)
	return names
}

func IsFile(path string) bool {
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

func IsDir(filename string) bool {
	fd, err := os.Stat(filename)
	if err != nil {
		return false
	}
	fm := fd.Mode()
	return fm.IsDir()
}

func WriteFile(path string, data string) error {
	err := Mkdir(path)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path, []byte(data), 0666)
	return err
}

func Mkdir(path string) error {
	dir := filepath.Dir(path)
	err := os.MkdirAll(dir, 0777)
	return err
}

func CopyFile(dstName string, srcName string) (int64, error) {
	err := Mkdir(dstName)
	if err != nil {
		return 0, err
	}
	src, err := os.Open(srcName)
	if err != nil {
		return 0, err
	}
	defer src.Close()
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return 0, err
	}
	defer dst.Close()
	return io.Copy(dst, src)
}

func CopyDir(dstName string, srcName string) error {
	err := filepath.Walk(srcName, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		base, _ := filepath.Rel(srcName, path)
		dst := filepath.Join(dstName, base)
		_, err = CopyFile(dst, path)
		return err
	})
	return err
}
