// string project string.go
package strings

import (
	"fmt"
	"strconv"
)

func Str_Method() {
	bs := make([]byte, 0, 100) //创建一个slice
	fmt.Println(bs)
	bs = strconv.AppendInt(bs, 4567, 10)
	fmt.Println(bs)

	bt := []byte("123456")
	fmt.Println(bt)

	bs = strconv.AppendBool(bs, false)
	bs = strconv.AppendQuote(bs, "abcxx")
	bs = strconv.AppendQuoteRune(bs, '我')
	fmt.Println(bs)
	fmt.Println(string(bs))

	fmt.Println(strconv.FormatInt(1234, 2))
	fmt.Println(strconv.ParseInt("10011010010", 2, 0))

}
