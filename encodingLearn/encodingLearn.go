// encodingLearn project encodingLearn.go
package encodingLearn

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func Jiangyou() {
	println("打酱油~")
}

func Ecoding_JSON() {
	buffer := bytes.NewBuffer(nil)
	encoder := json.NewEncoder(buffer)
	decoder := json.NewDecoder(buffer)

	data := map[string]interface{}{
		"s": "Hello zhangql",
		"i": 199999,
		"m": map[string]int{"TEST": 1, "GOLANG": 2},
	}

	//fmt.Print(json.Marshal(data))

	encoder.Encode(data)
	fmt.Printf("%#v\n", buffer.String())

	d := make(map[string]interface{})
	decoder.Decode(&d)
	fmt.Printf("%#v\n", d)
}

type P struct{ X, Y int }

type Data struct {
	S  string
	I  int
	T  struct{ A, B string }
	M  map[string]string
	M2 map[string]interface{}
	E  []int
	P
	f float32
}

func Ecoding_JSON_Map() {
	data := Data{
		S:  "Hello world",
		I:  999,
		T:  struct{ A, B string }{"Hello", "zhangql"},
		M:  map[string]string{"abc": "ABC", "def": "DEF"},
		M2: map[string]interface{}{"HIG": 888, "JKL": "M2"},
		E:  []int{1, 2, 3, 4},
		P:  P{777, 666},
		f:  1.12,
	}

	j, _ := json.MarshalIndent(&data, "", "  ")
	fmt.Printf("%s\n", string(j))

	var d Data
	json.Unmarshal(j, &d)
	fmt.Printf("%+v\n", d)
}
