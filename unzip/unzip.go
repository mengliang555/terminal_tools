package main

import (
	"fmt"
	"github.com/alexmullins/zip"
	"io"
	"os"
)

func unzip(filepath string) {
	// 打开.7z文件
	archive, err := zip.OpenReader(filepath)
	if err != nil {
		fmt.Println("无法打开.7z文件:", err)
		return
	}
	defer archive.Close()

	// 遍历.7z文件中的所有文件和文件夹
	for _, file := range archive.File {
		// 打开文件
		srcFile, err := file.Open()
		if err != nil {
			fmt.Println("无法打开文件:", err)
			return
		}
		defer srcFile.Close()

		// 创建目标文件
		dstFile, err := os.Create(file.Name)
		if err != nil {
			fmt.Println("无法创建目标文件:", err)
			return
		}
		defer dstFile.Close()

		// 将源文件内容复制到目标文件
		_, err = io.Copy(dstFile, srcFile)
		if err != nil {
			fmt.Println("无法复制文件内容:", err)
			return
		}
	}

	fmt.Println("文件解压完成")
}
