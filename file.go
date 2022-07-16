package utils

import (
	"bufio"
	"io"
	"net/http"
	"os"
	"path"
)

//SaveFile 读取文件流保存到本地
func SaveFile(rd io.Reader, destPath string, fileName string) (*os.File, error) {
	destPath = MakePath(destPath, 0766)

	file, err := os.Create(destPath + fileName)
	if err != nil {
		return nil, err
	}
	// 获得文件的writer对象
	writer := bufio.NewWriter(file)
	io.Copy(writer, rd)
	file.Seek(0, 0)
	return file, nil
}

//FetchFile 拉取远程图片
func FetchFile(fileUrl string, destPath string, fileName string) (*os.File, error) {
	destPath = MakePath(destPath, 0766)
	//imgUrl := "https://cdn2.jianshu.io/assets/default_avatar/9-cceda3cf5072bcdd77e8ca4f21c40998.jpg"

	if len(fileName) < 1 {
		fileName = path.Base(fileUrl)
	}
	res, err := http.Get(fileUrl)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	// 获得get请求响应的reader对象
	//fmt.Printf("网络文件大小为[%v]:", res.ContentLength)
	reader := bufio.NewReaderSize(res.Body, int(res.ContentLength)+1024)

	file, err := os.Create(destPath + fileName)
	if err != nil {
		return nil, err
	}
	// 获得文件的writer对象
	writer := bufio.NewWriter(file)
	io.Copy(writer, reader)
	file.Seek(0, 0)
	return file, nil
}

//PathExists 文件或文件夹是否存在
func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

//MakePath 创建目录
func MakePath(path string, perm os.FileMode) string {
	if res := PathExists(path); !res {
		err := os.MkdirAll(path, perm)
		if err != nil {
			panic(err)
		}
	}
	return path
}
