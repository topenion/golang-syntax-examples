package main

import "fmt"

const (
	byt  = iota + 8
	kbyt = 1 << (iota * 10)
	mbyt
	gbyt
)

func main() {
	fmt.Println(byt, kbyt, mbyt, gbyt)
}
