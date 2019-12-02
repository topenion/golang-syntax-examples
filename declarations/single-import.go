package main

import "fmt" //single Line Import

var zeroValueTestInt int //single line declaration

func singleLineImport() {
	message := "Single Line Assignment and declarations only work inside blocks"

	fmt.Println(message) //Print message with new line

	fmt.Printf("Zero Value of %T is %v\n", zeroValueTestInt, zeroValueTestInt) //Formatted Printing %T is type %v is value
}
