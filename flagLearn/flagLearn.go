// flagLearn project flagLearn.go
package flagLearn

import (
	"flag"
	"fmt"
	"os"
)

func Flag_Method1() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "MyProgram, version 0.0.0.1\n")
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}

	//接收参数方法 参数名 默认值 帮助信息中的解释
	file := flag.String("file", "default.txt", "input filename")
	count := flag.Int("Count", 0, "process count.")

	flag.Parse()
	fmt.Printf("Flag: %d, %s, %d\n", flag.NFlag(), *file, *count)
	fmt.Printf("Args: %d, %v\n", flag.NArg(), flag.Args())

	flag.VisitAll(func(f *flag.Flag) {
		fmt.Printf("%s = %v (default:%v)\n", f.Name, f.Value, f.DefValue)
	})

	flag.Visit(func(f *flag.Flag) {
		fmt.Printf("%v\n", f.Value)
	})

	f := flag.Lookup("Count")
	fmt.Printf("%s\n", f.Value)
}

func Jiangyou() {
	println("Flag Learn 打酱油~")
}
