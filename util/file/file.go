package file

import (
	"fmt"
	"math/rand"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

func AbsPath(fpath string) (string, error)  {
	isAbs := path.IsAbs(fpath)
	if isAbs {
		return fpath, nil
	}
	// fpath如果是相对路径，filepath.Abs是基于可执行文件位置来计算绝对路径的
	fpath, err := filepath.Abs(fpath)
	if err != nil {
		return "", err
	}
	return fpath, nil
}

func AbsolutePath(filePath string) string {
	_, filename, _, _ := runtime.Caller(0)
	absolutePath := path.Join(path.Dir(filename), filePath)
	return absolutePath
}

// 随机数范围： [min, max)
func RandInt(min, max int) int {
	numRange := max - min
	ret :=  rand.Intn(numRange) + min
	fmt.Println("rand ret:-->", ret)
	return ret
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func NewFile() {
	filepath.
}
