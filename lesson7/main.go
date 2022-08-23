package main

import (
	"fmt"
	"io"
)

type dataByte struct {
	read string
	done bool
}

func main() {
	var a = &dataByte{}
	_, err := io.WriteString(a, "Привет")
	if err != nil {
		fmt.Println(err)
		return
	}
	s, err := io.ReadAll(a)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(s))
}

func (md *dataByte) Write(p []byte) (n int, err error) {
	md.read = string(p)
	return len(p), nil
}

func (md *dataByte) Read(p []byte) (n int, err error) {
	if md.done {
		return 0, io.EOF
	}
	for i, b := range []byte(md.read) {
		p[i] = b
	}
	md.done = true
	return len(md.read), nil
}
