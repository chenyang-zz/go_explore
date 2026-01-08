package io

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

// 创建文件
func CreateFile(filename string) {
	os.Remove(filename)
	if file, err := os.Create(filename); err != nil {
		fmt.Printf("create file failed: %s\n", err)
	} else {
		defer file.Close()
		file.Chmod(0o666)
		fmt.Printf("fd=%d\n", file.Fd())
		file.WriteString("多情应笑我\n")
		info, _ := file.Stat()
		fmt.Printf("is dir %t\n", info.IsDir())
		fmt.Printf("modify time: %s\n", info.ModTime())
		fmt.Printf("mode %v\n", info.Mode())
		fmt.Printf("file name %s\n", info.Name())
		fmt.Printf("size %dB\n", info.Size())
	}

	os.Mkdir("../data/sys", os.ModePerm)
	os.MkdirAll("../data/sys/a/b/c", os.ModePerm)

	os.Rename("../data/sys/a", "../data/sys/p")
	os.Rename("../data/sys/p/b/c", "../data/sys/p/c")

	// os.Remove("../data/sys/p/b")
	// os.RemoveAll("../data/sys")
}

// 遍历一个目录
func WalkDir(path string) error {
	filepath.Walk(path, func(subpath string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.Mode().IsDir() && path != subpath {
			fmt.Printf("path is dir %s\n", subpath)
		} else if info.Mode().IsRegular() {
			fmt.Printf("path is file %s basename %s", subpath, info.Name())
		}

		return nil
	})

	return nil
}
