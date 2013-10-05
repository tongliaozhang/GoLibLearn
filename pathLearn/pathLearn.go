// pathLearn project pathLearn.go
package pathLearn

import (
	"fmt"
	"os"
	"path/filepath"
)

func PathLearn() {
	fmt.Println(filepath.Abs("test2.dat"))
	fmt.Println(os.ExpandEnv("$JAVA_HOME/demo.txt"))
	fmt.Println(os.ExpandEnv("$HOME" + "~/demo.txt"[1:]))

	fmt.Println(filepath.IsAbs("c:/etc/a.conf"), filepath.IsAbs("./demo.txt"))

	f, _ := filepath.Abs("./test.dat")

	println(filepath.Dir(f))
	println(filepath.Base(f))
	println(filepath.Ext(f))
}
