package main

//multiple import can be clubbed
import (
	"fmt"
	"math"
)

//multiple vars can be clubbed
var (
	clubbedString            string
	clubedBool, clubbedFloat = false, 9.0 //single line
)

//multiple constants can be clubbed

const (
	clubbedConstA = "A"
	clubbedConstB = "B"
)

func withParamsAndReturnValue(param1 bool, param2 string) (int, error) {
	clubedBool = param1
	clubbedString = param2
	// clubbedConstA =  10 //wont work as const can't be reassigned

	//functions can return multiple value
	num, err := fmt.Println(math.Sqrt(clubbedFloat), clubedBool, clubbedString, clubbedConstA, clubbedConstB)
	return num, err
}
