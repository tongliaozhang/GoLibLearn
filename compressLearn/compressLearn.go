// compressLearn project compressLearn.go
package compressLearn

import (
	"archive/zip"
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func ZlibCompress() {
	buffer := bytes.NewBuffer(nil)

	w := zlib.NewWriter(buffer)

	io.WriteString(w, "hello world")

	w.Close()

	fmt.Printf("% x\n", buffer.Bytes())

	ioutil.WriteFile("test2.dat", buffer.Bytes(), 0644)
}

func ZlibUnCompress() (flag bool, err error) {
	defer func() {
		e := recover()
		if e == nil {
			return
		}
		panic("文件不存在！")
	}()

	data, err := ioutil.ReadFile("test2.dat")

	r, _ := zlib.NewReader(bytes.NewBuffer(data))

	bs, _ := ioutil.ReadAll(r)

	println(string(bs))
	flag = true
	return flag, err
}

func ZipFile() {
	file, _ := os.Create("test.zip")
	defer file.Close()

	writer := zip.NewWriter(file)
	defer writer.Close()

	data, _ := ioutil.ReadFile("test2.dat")

	w, _ := writer.Create("test2.dat")

	w.Write(data)
}

func ZipFileAndInfo() {
	file, _ := os.Create("test1.zip")
	defer file.Close()
	writer := zip.NewWriter(file)
	defer writer.Close()
	for _, name := range []string{"test2.dat"} {
		info, _ := os.Stat(name)
		header, _ := zip.FileInfoHeader(info)
		w, _ := writer.CreateHeader(header)
		body, _ := ioutil.ReadFile(name)
		w.Write(body)
	}
}

func UnZipFile() {
	zr, _ := zip.OpenReader("test.zip")
	defer zr.Close()

	for _, f := range zr.File {
		br, _ := f.Open()
		bs, _ := ioutil.ReadAll(br)

		fmt.Printf("%s:%s\n", f.Name, string(bs))
	}
}

func Jiangyou() {
	println("打酱油~")
}
