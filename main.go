// GoLibLearn project main.go
package main

import (
	//"GoLibLearn/compressLearn"
	//"GoLibLearn/cryptoLearn"
	//"GoLibLearn/encodingLearn"
	//"GoLibLearn/flagLearn"
	//"GoLibLearn/ioLearn"
	//"GoLibLearn/pathLearn"
	//"GoLibLearn/net"
	//"runtime"
	//"GoLibLearn/os"
	"GoLibLearn/strings"
	//"fmt"
)

func main() {
	//ioLearn.OpenFile()
	//ioLearn.WriteBinary()
	//ioLearn.ReadBinary()
	//pathLearn.PathLearn()
	//runtime.GOMAXPROCS(runtime.NumCPU())
	//go net.Server()
	//go net.Client()
	//select {}

	//os.SystemCmd()
	//os.Process()
	//os.Process_Output()

	strings.Str_Method()

	//fmt.Println("Hello World!")
	//compressLearn.Jiangyou()
	//cryptoLearn.Jiangyou()
	//encodingLearn.Jiangyou()
	//flagLearn.Jiangyou()

	//flagLearn.Flag_Method1()

	/*compressLearn.ZlibCompress()

	flag, err := compressLearn.ZlibUnCompress()

	if err != nil {
		println(flag)
	}
	println(flag)
	*/

	//compressLearn.ZipFile()
	//compressLearn.UnZipFile()
	//compressLearn.ZipFileAndInfo()

	//cryptoLearn.Md5_crypto()
	//cryptoLearn.Sha256_crypto()
	//cryptoLearn.Sha512_crypto()
	//cryptoLearn.HMAC_crypto()
	//cryptoLearn.Rand_crypto()
	//cryptoLearn.DES_crypto()
	//cryptoLearn.RAS_crypto()
	//cryptoLearn.Key_crypto()
	//cryptoLearn.Check_Crypto()

	//encodingLearn.Ecoding_JSON()
	//encodingLearn.Ecoding_JSON_Map()

}

/*
recover 和 panic 例子
func readAll(r io.Reader, capacity int64) (b []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0, capacity))
	// If the buffer overflows, we will get bytes.ErrTooLarge.
	// Return that as an error. Any other panic remains.
	defer func() {
		e := recover()
		if e == nil {
			return
		}
		if panicErr, ok := e.(error); ok && panicErr == bytes.ErrTooLarge {
			err = panicErr
		} else {
			panic(e)
		}
	}()
	_, err = buf.ReadFrom(r)
	return buf.Bytes(), err
}
*/
