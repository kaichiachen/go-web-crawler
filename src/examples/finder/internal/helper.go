package internal

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// checkDirPath 会检查目录路径。
func checkDirPath(dirPath string) (absDirPath string, err error) {
	if dirPath == "" {
		err = fmt.Errorf("invalid dir path: %s", dirPath)
		return
	}
	if filepath.IsAbs(dirPath) {
		absDirPath = dirPath
	} else {
		absDirPath, err = filepath.Abs(dirPath)
		if err != nil {
			return
		}
	}
	var dir *os.File
	dir, err = os.Open(absDirPath)
	if err != nil && !os.IsNotExist(err) {
		return
	}
	if dir == nil {
		err = os.MkdirAll(absDirPath, 0700)
		if err != nil && !os.IsExist(err) {
			return
		}
	} else {
		var fileInfo os.FileInfo
		fileInfo, err = dir.Stat()
		if err != nil {
			return
		}
		if !fileInfo.IsDir() {
			err = fmt.Errorf("not directory: %s", absDirPath)
			return
		}
	}
	return
}

// Record 用于记录日志。
func Record(level byte, content string) {
	if content == "" {
		return
	}
	switch level {
	case 0:
		log.Println(content)
	case 1:
		log.Println(content)
	case 2:
		log.Println(content)
	}
}
