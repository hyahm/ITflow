package gapublic

import (
	// "bufio"
	"errors"
	// "fmt"

	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func ReadAll(filePath string) (string, error) {
	if !checkFileIsExist(filePath) {
		return "", errors.New("file not found")
	}
	f, err := os.Open(filePath)
	defer f.Close()
	if err != nil {
		return "", err
	}
	s, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}
	return string(s), nil
}

func WriteAll(filePath string, content string) error {
	if !checkFileIsExist(filePath) {
		_, _ = os.Create(filePath)
	}
	f, _ := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	defer f.Close()
	f_content := strings.Join([]string{content}, "")
	buf := []byte(f_content)
	f.Write(buf)

	return nil
}

// 检查是否可以创建文件夹
func CheckCreatePath(path string, perm os.FileMode) {
	err := os.MkdirAll(path, perm)
	if err != nil {
		panic(err)
	}
}

func JoinPathAndFile(path string, filename string) string {
	if filename == "" {
		return ""
	}

	return filepath.Join(path, filename)
}
