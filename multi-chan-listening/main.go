package main

import "fmt"

var q = make(chan bool)

func main() {
	str1 := "helloworld"
	str2 := "HELLOWORLD"

	f1 := func() chan rune {
		passer := make(chan rune)
		go func() {
			for _, char := range str1 {
				passer <- char
			}
			close(passer)
		}()

		return passer
	}

	f2 := func() chan rune {
		passer := make(chan rune)
		go func() {
			for _, char := range str2 {
				passer <- char
			}
			close(passer)
			q <- false
		}()
		return passer
	}
	s1 := f1()
	s2 := f2()
	chanreceiver(s1, s2)

}

func chanreceiver(c1, c2 chan rune) {
	var str3 string
	var h rune
	for {
		select {
		case h = <-c1:
			str3 = str3 + string(h)
		case h = <-c2:
			str3 = str3 + string(h)
		case <-q:
			fmt.Println(str3)
			return
		}
	}
}
