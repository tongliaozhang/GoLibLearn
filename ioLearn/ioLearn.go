// ioLearn project ioLearn.go
package ioLearn

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"os"
)

func OpenFile() {
	f, _ := os.Open("text.txt")
	defer f.Close()

	r := bufio.NewReaderSize(f, 4096)

	for {
		line, isPerfix, err := r.ReadLine()

		if isPerfix {
			print(string(line))
		} else if len(line) > 0 {
			println(string(line))
		}

		if err == io.EOF {
			break
		}
	}
}

func checkError(err interface{}) {
	if err != nil {
		log.Fatal(err)
	}
}

func WriteBinary() {
	fmt.Println("************************************")
	f, err := os.Create("test2.dat")
	checkError(err)

	defer func() {
		f.Sync()
		f.Close()
	}()

	var i int32 = 0x1234
	checkError(binary.Write(f, binary.LittleEndian, i))

	var d float32 = 0.1234
	checkError(binary.Write(f, binary.LittleEndian, d))

	var s string = "Hello ioUtil      "
	checkError(binary.Write(f, binary.LittleEndian, int32(len(s))))
	_, err = f.WriteString(s)

	checkError(err)

	buffer := bytes.NewBuffer(nil)

	var ii int32 = 0x1234
	binary.Write(buffer, binary.LittleEndian, ii)

	buffer.WriteString("12312221314")

	fmt.Printf("% x\n", buffer.Bytes())
}

func ReadBinary() {
	f, err := os.Open("test2.dat")
	checkError(err)
	defer f.Close()

	var i int32
	checkError(binary.Read(f, binary.LittleEndian, &i))
	var d float32
	checkError(binary.Read(f, binary.LittleEndian, &d))
	var l int32
	checkError(binary.Read(f, binary.LittleEndian, &l))
	s := make([]byte, l)
	_, err = f.Read(s)
	fmt.Printf("%#x; %f; %s;\n", i, d, string(s))
}
