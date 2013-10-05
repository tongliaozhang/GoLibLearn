// os project os.go
package os

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"time"
)

func SystemCmd() {
	//os.MkdirAll("./test/test1/test2", 0777)

	if err := os.Mkdir("dir", 0755); os.IsExist(err) {
		println("dir is exist !")
	}

	if file, err := os.Open("abc.txt"); os.IsNotExist(err) {
		println("file not exist")
	} else {
		file.Close()
	}

	file, err := os.Create("abc.txt")
	if err == nil {
		println(file.Name(), "create success")
	}

	os.Link("./a.txt", "./b.txt")
	sa, _ := os.Stat("./a.txt")
	sb, _ := os.Stat("./b.txt")
	fmt.Println(os.SameFile(sa, sb))
	os.Symlink("./a.txt", "./c.txt")
	//fmt.Println(os.Readlink("./c.txt"))

	envs := os.Environ()
	for _, s := range envs {
		fmt.Println(s)
	}

}

func logErr(err interface{}) {
	if err != nil {
		log.Fatal(err)
	}
}

func Process() {
	//cmd := exec.Command("dir", "-p", "./")
	cmd := exec.Command("ls", "-altr", "./")

	stdout, _ := cmd.StdoutPipe()

	logErr(cmd.Start())

	fmt.Println(cmd.Process.Pid)

	bs, _ := ioutil.ReadAll(stdout)
	fmt.Println(string(bs))

	time.Sleep(time.Minute)

	logErr(cmd.Wait())
}

func Process_Output() {
	bs, err := exec.Command("ls", "-altr", "./").Output()

	logErr(err)

	fmt.Println(string(bs))
}
