//Package Main aims at multiple type of declarations you can do in golang
package main

//Single Import
import "fmt"

//single vars
var uninitiallized int
var initialized = true //type is inferred

//single const
const constantString = "I am Constant"

func simpleFunction() {
	shortDeclaration := "Can not be used outside {}"
	fmt.Println("Single Line Values are", uninitiallized, initialized, constantString, shortDeclaration)
}
func main() {
	simpleFunction()

}
