package zip

import (
	"archive/zip"
	"io"
	"log"
	"os"
)

//Zip 压缩
func Zip(zipFile string, fileList []string) error {
	// 创建 zip 包文件
	fw, err := os.Create(zipFile)
	if err != nil {
		log.Fatal()
	}
	defer fw.Close()

	// 实例化新的 zip.Writer
	zw := zip.NewWriter(fw)
	defer func() {
		// 检测一下是否成功关闭
		if err := zw.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	for _, fileName := range fileList {
		fr, err := os.Open(fileName)
		if err != nil {
			return err
		}
		fi, err := fr.Stat()
		if err != nil {
			return err
		}
		// 写入文件的头信息
		fh, err := zip.FileInfoHeader(fi)
		w, err := zw.CreateHeader(fh)
		if err != nil {
			return err
		}
		// 写入文件内容
		_, err = io.Copy(w, fr)
		if err != nil {
			return err
		}
	}
	return nil
}

//Unzip 解压缩
func Unzip(zipFile string) error {
	zr, err := zip.OpenReader(zipFile)
	defer zr.Close()
	if err != nil {
		return err
	}

	for _, file := range zr.File {
		// 如果是目录，则创建目录
		if file.FileInfo().IsDir() {
			if err = os.MkdirAll(file.Name, file.Mode()); err != nil {
				return err
			}
			continue
		}
		// 获取到 Reader
		fr, err := file.Open()
		if err != nil {
			return err
		}

		fw, err := os.OpenFile(file.Name, os.O_CREATE|os.O_RDWR|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}
		_, err = io.Copy(fw, fr)
		if err != nil {
			return err
		}
		fw.Close()
		fr.Close()
	}
	return nil
}
