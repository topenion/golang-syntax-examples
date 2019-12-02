package main

import "fmt"

var (
	str   string
	i     int
	i64   int64
	ui64  uint64
	byt   byte
	boo   bool
	cn64  complex64
	cn128 complex128
	ru    rune
)

func printZeroValues() { //not exportable : function name starts with small letters

	fmt.Printf("Zero Value of %T is %q\n", str, str) //%q for quoted
	fmt.Printf("Zero Value of %T is %v\n", i64, i64)
	fmt.Printf("Zero Value of %T is %v\n", ui64, ui64)
	fmt.Printf("Zero Value of %T is %v\n", byt, byt)
	fmt.Printf("Zero Value of %T is %v\n", boo, boo)
	fmt.Printf("Zero Value of %T is %v\n", cn64, cn64)
	fmt.Printf("Zero Value of %T is %v\n", cn128, cn128)
	fmt.Printf("Zero Value of %T is %v\n", ru, ru)
}
