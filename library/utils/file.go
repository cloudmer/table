package utils

import (
	"os"
	"table/share"
)

// 创建文件夹 没有则 创建
func GenerateDir(dirPath string) (path string, err error) {
	// 检查 文件夹是否存在
	_, err = os.Stat(dirPath)
	if err == nil {
		return dirPath, nil
	}
	// 文件夹存在
	if !os.IsNotExist(err) {
		return dirPath, nil
	}

	// 文件夹不存在 则创建文件夹
	err = os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return "", err
	}
	// 文件夹创建成功
	return dirPath, nil
}

// 生成文件
func GenerateFile(fileName string) (path string, err error) {
	// 文件不存在 创建文件
	share.ShareLoggerFile, err = os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		return "", err
	}
	return fileName, nil
}

// 文件内容写入
func FileWrite(file *os.File, contents string) error {
	_, err := file.WriteString(contents)
	return err
}
