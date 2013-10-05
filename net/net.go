// net project net.go
package net

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func Server() {
	listen, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, _ := listen.Accept()

		go func() {
			defer conn.Close()
			rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))

			for {
				s, err := rw.ReadString('\n')
				if err == io.EOF {
					fmt.Println("[server] client close ......")
					return
				}

				println("[server]", s)
				fmt.Fprintln(rw, "server: %s\n", s)
				rw.Flush()
			}
		}()
	}
}

func Client() {
	conn, _ := net.Dial("tcp", "localhost:8080")
	defer conn.Close()

	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))

	fmt.Fprintln(rw, "hello 10.1")
	rw.Flush()
	s, _ := rw.ReadString('\n')
	fmt.Println("[client]", s)
}
