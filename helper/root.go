// Package helper
// @Description:
package helper

import (
	"errors"
	"log"
	"os"
)

var (
	MakeDirError = errors.New("File or folder creation failed")
)

/**
 *  FileIsExist
 *  @Description:检查文件是否存在，同时返回文件详情
 *  @param filename
 *  @return os.FileInfo
 *  @return bool
 **/
func FileOrDirIsExist(name string) (os.FileInfo, bool) {
	info, err := os.Stat(name)
	if err == nil {
		return info, true
	}
	if os.IsNotExist(err) {
		return nil, false
	}
	return info, true
}

/**
 *  MakeDir
 *  @Description: 递归创建文件夹
 *  @param name
 *  @return bool
 **/
func MakeDir(path string) bool {
	if errs := os.MkdirAll(path, os.ModePerm); errs != nil {
		log.Fatal(MakeDirError, errs)
		return false
	}
	return true
}

/**
 *  CreateFile
 *  @Description: 创建文件
 *  @param name
 *  @return bool
 **/
func CreateFile(name string) bool {
	f, err := os.Create(name)
	if err != nil {
		return false
	}
	defer f.Close()
	return true
}
