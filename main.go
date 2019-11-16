/*This is main package. All Files in same directory must have package main or package main_test
* Also there must be a main function in any of the main package files, which is the starting
* point of the program.
* NOTE: package main makes executables, other packages are just packages to be used in your program*/
package main

//multiline import
import (
	"fmt"
	"github.com/topenion/golang-syntax-examples/declarations"
)

func main() {
	fmt.Println("The Main function")
	declarations.RunDeclarations()
}
